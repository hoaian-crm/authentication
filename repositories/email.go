package repositories

import (
	"main/base"
	"main/config"
	"main/models"
	"main/utils"
)

type EmailRepository struct {
	base.Repository[models.Email]
}

func (EmailRepository EmailRepository) SendMailToUser(data *models.Email) {

	userRepository := UserRepository{}

	userId, _ := utils.StringToNumber(data.SendTo)

	user, _ := userRepository.FindOne(&models.User{
		BaseModel: models.BaseModel{
			ID: uint(userId),
		},
	})

	config.MailSender.SendMail(config.MailData{
		To:      user.Email,
		Subject: data.Subject,
		Content: data.Content,
	})
}
