package services

import (
	"main/base"
	"main/config"
	"main/constants"
	user_dto "main/dtos/user"
	"main/models"
	"main/prototypes/gen/go/pb"
	"main/repositories"
	"main/responses"
	"main/utils"

	c "context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserService struct {
	base.Service[models.User]
	cache repositories.CacheRepository[models.User]
}

func (userService UserService) Register(context *gin.Context) {

	data := context.MustGet("data").(models.User)
	db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)

	data.Role = models.Role{
		BaseModel: models.BaseModel{
			ID: config.EnvirontmentVariables.DefaultRoleId,
		},
	}

	db.Create(&data)

	response := config.Response{
		Data: config.ResponseData{
			Limit:  1,
			Total:  1,
			Offset: 0,
			Result: data,
		},
		Messages: []config.Message{},
	}

	template := "otp"

	config.GrpcClient.Emit(c.Background(), &pb.EmitEventDto{
		Name: "user_register",
		Payload: &pb.EmitEventDto_Payload{
			Value: &pb.EmitEventDto_Payload_Mail{
				Mail: &pb.SendMailDto{
					To:       "",
					Subject:  "",
					Template: &template,
					Context: map[string]string{
						"otp": data.OtpCode,
					},
				},
			},
		},
	})

	response.Created(context)
}

func (userService UserService) Login(context *gin.Context) {

	data := context.MustGet("data").(user_dto.LoginDto)
	db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)

	userRecord := models.User{
		Email:  data.Email,
		Active: true,
	}

	db.Where(&userRecord).Preload("Role").First(&userRecord)

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

func (userService UserService) GetProfile(context *gin.Context) {

	db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)

	userId := context.MustGet("UserId").(int64)

	var user = models.User{
		BaseModel: models.BaseModel{
			ID: userId,
		},
	}

	db.Where(&user).Preload("Role").Preload("Role.Permissions").First(&user)

	response := config.Response{
		Data: config.ResponseData{
			Result: user,
		},
		Messages: []config.Message{config.Messages["get_success"]},
	}
	response.GetSuccess(context)
}

func (userService UserService) ActiveUser(context *gin.Context) {
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

func (userService UserService) UpdatePassword(context *gin.Context) {

	userId := context.MustGet("userId").(int64)

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

	// Prepare data
	query := context.MustGet("query").(user_dto.ListUserDto)
	query.SetDefaults()
	db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)

	// Search
	db.Where("\"displayName\" like ? or email like ?", "%"+query.Keyword+"%", "%"+query.Keyword+"%")

	// Order
	if query.Order != "" {
		db.Order(clause.OrderByColumn{Column: clause.Column{Name: query.Order}, Desc: query.Direction == "desc"})
	}

	// Count records
	var total int64
	db.Count(&total)
	db.Limit(query.Limit)

	// Offset
	db.Offset(query.Offset)

	// Get result
	result := []models.User{}
	db.Preload("Role").Find(&result)

	// Response to client
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
