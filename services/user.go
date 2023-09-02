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

func (userService UserService) Login(context *gin.Context) {

	data := base.GetData[user_dto.LoginDto](context)

	userRecord, _ := userService.Repository.FindOne(&models.User{
		Email:  data.Email,
		Active: true,
	})

	if utils.ComparePassword(userRecord.Password, data.Password) {
		messages := []config.Message{config.Messages["login_success"]}
		data := config.ResponseData{
			Result: responses.LoginReponse{
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

	user := context.MustGet("user").(models.User)

	response := config.Response{
		Data: config.ResponseData{
			Result: user,
		},
		Messages: []config.Message{config.Messages["get_success"]},
	}
	response.GetSuccess(context)
}

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
