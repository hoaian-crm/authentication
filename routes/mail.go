package routes

import (
	"main/services"

	"github.com/gin-gonic/gin"
)

func NewMailController(route *gin.RouterGroup) {

	emailService := services.EmailService{}

	route = route.Group("/mails")
	{
		route.POST("/send_mail_verify", func(context *gin.Context) {
			emailService.SendVerifyMail(context)
		})
	}
}
