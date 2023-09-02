package services

import (
	"main/base"
	"main/config"
	user_dto "main/dtos/user"
	"main/models"
	"main/repositories"
	"main/responses"
	"main/utils"

	"github.com/gin-gonic/gin"
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
	userService.Service.CreateOne(context)
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

	data := base.GetData[user_dto.LoginDto](context)

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

	user := context.MustGet("user").(models.User)

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
	data := base.GetData[user_dto.ActiveUserDto](context)
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
	data := base.GetData[user_dto.UpdatePassword](context)
	userToken := context.MustGet("user").(models.User)

	user, _ := userService.Repository.FindOne(&models.User{
		ID: userToken.ID,
	})

	ok := utils.ComparePassword(user.Password, data.CurrentPassword)
	if !ok {
		response := config.Response{
			Data:     config.NoData(),
			Messages: []config.Message{config.Messages["invalid_email_password"]},
		}
		response.BadRequest(context)
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
