package services

import (
	"main/base"
	"main/config"
	"main/constants"
	email_dto "main/dtos/email"
	user_dto "main/dtos/user"
	"main/models"
	mail_queue "main/queue/mail"
	"main/repositories"
	"main/responses"
	"main/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	base.Service[models.User]
	cache repositories.CacheRepository[models.User]
}

// @Summary	Register new user
// @Schemes
// @Description				Call this api to register new user
// @Accept						json
// @Produce					json
// @Router						/users/register [post]
// @Tags User
// @success 200 {object} config.Response{data=responses.RegisterResponse}
// @Param   request  body     user_dto.Create  true "Login data"
func (userService UserService) Register(context *gin.Context) {
	var dto models.User
	if err := context.ShouldBind(&dto); err != nil {
		messages := config.MessagesBuilder(err)
		response := config.Response{
			Data:     config.NoData(),
			Messages: messages,
		}
		response.BadRequest(context)
		return
	}

	dto.Iss = config.EnvirontmentVariables.UserIss

	createdRecord := userService.Repository.CreateOne(&dto)
	if createdRecord.Error != nil {
		response := config.Response{
			Data: config.NoData(),
			Messages: []config.Message{{
				Code:        -1,
				Description: createdRecord.Error.Error(),
			}},
		}
		response.InternalServerError(context)
		return
	}
	response := config.Response{
		Data: config.ResponseData{
			Limit:  1,
			Total:  1,
			Offset: 0,
			Result: dto,
		},
		Messages: []config.Message{},
	}

	go mail_queue.SendMailToUser(email_dto.SendMailToUserDto{
		SendTo:  dto.Email,
		Content: dto.OtpCode,
		Subject: "Friendify Verify Code",
	})

	response.Created(context)
}

// @Summary	Login user
// @Schemes
// @Description				Call this api to login
// @Accept						json
// @Produce					json
// @Router						/users/login [post]
// @Tags User
// @success 200 {object} config.Response{data=responses.LoginResponse}
// @Param   request  body     user_dto.LoginDto  true "Login data"
func (userService UserService) Login(context *gin.Context) {

	// data := base.GetData[user_dto.LoginDto](context)
	data := context.MustGet("data").(user_dto.LoginDto)

	userRecord, _ := userService.Repository.FindOne(&models.User{
		Email:  data.Email,
		Active: true,
	})

	if utils.ComparePassword(userRecord.Password, data.Password) {
		messages := []config.Message{config.Messages["login_success"]}
		data := config.ResponseData{
			Result: responses.LoginResponse{
				AccessToken: userRecord.SignToken(),
			},
		}
		response := config.Response{
			Messages: messages,
			Data:     data,
		}
		response.Created(context)
		return
	}

	response := config.Response{
		Messages: []config.Message{config.Messages["invalid_email_password"]},
		Data:     config.NoData(),
	}

	response.BadRequest(context)
}

// @Tags User
// @Summary Get user's profile
// @Description  Get user's profile by access token
// @Accept json
// @Schemes
// @Router /users/profile [get]
// @Security BearerAuth
func (userService UserService) GetProfile(context *gin.Context) {

	userRepository := repositories.UserRepository{}

	userId := context.MustGet("userId").(uint)

	user, _ := userRepository.FindOne(&models.User{
		BaseModel: models.BaseModel{
			ID: userId,
		},
	})

	response := config.Response{
		Data: config.ResponseData{
			Result: user,
		},
		Messages: []config.Message{config.Messages["get_success"]},
	}
	response.GetSuccess(context)
}

// @Tags User
// @Summary Active account
// @Description  Active new user
// @Accept json
// @Schemes
// @Router /users/active [put]
// @Param request body user_dto.ActiveUserDto true "Active user data"
func (userService UserService) ActiveUser(context *gin.Context) {
	// data := base.GetData[user_dto.ActiveUserDto](context)
	data := context.MustGet("data").(user_dto.ActiveUserDto)
	userRecord, _ := userService.Repository.FindOne(&models.User{
		Email: data.Email,
	})
	success := userRecord.ActiveUser(data.OtpCode)
	if success {
		response := config.Response{
			Data: config.ResponseData{
				Result: userRecord,
			},
		}
		response.UpdateSuccess(context)
		return
	}
	response := config.Response{
		Data: config.NoData(),
		Messages: []config.Message{
			config.Messages["invalid_otp_code"],
		},
	}
	response.BadRequest(context)
}

// @Tags User
// @Summary Update password of user
// @Description  Call api to update password
// @Accept json
// @Schemes
// @Router /users/update_password [put]
// @Param request body user_dto.UpdatePassword true "Update password data"
// @Security BearerAuth
func (userService UserService) UpdatePassword(context *gin.Context) {

	userId := context.MustGet("userId").(uint)

	user, _ := userService.Repository.FindOne(&models.User{
		BaseModel: models.BaseModel{
			ID: userId,
		},
	})
	data := context.MustGet("data").(user_dto.UpdatePassword)

	ok := utils.ComparePassword(user.Password, data.CurrentPassword)
	if !ok {
		response := config.Response{
			Data:     config.NoData(),
			Messages: []config.Message{config.Messages["invalid_email_password"]},
		}
		response.BadRequest(context)
		return
	}

	user.Password = utils.HashPassword(data.NewPassword)
	config.Db.Save(&user)

	response := config.Response{
		Data: config.ResponseData{
			Result: user,
		},
		Messages: []config.Message{config.Messages["update_success"]},
	}
	response.UpdateSuccess(context)
}

func (userService UserService) ListUser(context *gin.Context) {
	query := context.MustGet("query").(user_dto.ListUserDto)
	query.SetDefaults()
	db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)

	db.Where("display_name like ? or email like ?", "%"+query.DisplayName+"%", "%"+query.Email+"%")

	var total int64
	db.Count(&total)
	db.Limit(query.Limit)

	result := []models.User{}
	db.Find(&result)

	response := config.Response{
		Data: config.ResponseData{
			Result: result,
			Total:  total,
			Offset: query.Offset,
			Limit:  query.Limit,
		},
		Messages: []config.Message{config.Messages["get_success"]},
	}

	response.GetSuccess(context)
}
