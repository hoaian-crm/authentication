package services

import (
	"main/base"
	"main/config"
	user_dto "main/dtos/user"
	"main/models"
	"main/repositories"
	"main/utils"

	"github.com/gin-gonic/gin"
)

type EmailService struct {
	base.Service[models.Email]
}

func (emailService EmailService) SendVerifyMail(context *gin.Context) {
	data := base.GetData[user_dto.SendVerifyMail](context)
	userService := UserService{}
	user, err := userService.Repository.FindOne(&models.User{
		Email: data.Email,
	})

	if err != nil {
		response := config.Response{
			Data:     config.NoData(),
			Messages: []config.Message{config.Messages["email_not_registered"]},
		}
		response.BadRequest(context)
	}

	emailRepostory := repositories.EmailRepository{}

	userIdString, _ := utils.NumberToString(int(user.ID))

	go emailRepostory.SendMailToUser(&models.Email{
		SendTo:  userIdString,
		Subject: "[Friendify] -> Verify account",
		Content: "Your otp code is: " + user.OtpCode,
	})

	response := config.Response{
		Data:     config.NoData(),
		Messages: []config.Message{},
	}

	response.Created(context)
}
