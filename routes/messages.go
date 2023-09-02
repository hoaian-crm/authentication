package routes

import (
	"main/config"
	messages_dto "main/dtos/messages"
	"main/middlewares"
	"main/models"
	"main/services"

	"github.com/gin-gonic/gin"
)

func NewMessageController(route *gin.RouterGroup) {
	config.Db.AutoMigrate(&models.MessageAbstract{})
	config.Db.AutoMigrate(&models.TextMessage{})

	messageService := services.MessageService{}
	channelService := services.ChannelService{}
	channelDetailService := services.ChannelDetailService{}

	route = route.Group("messages")
	{
		route.POST("/ask", // ? Api for chat with bot
			middlewares.Authorization(),
			middlewares.InitModel(&models.TextMessage{}),
			middlewares.BindBody[messages_dto.Create]("data"),
			channelService.UpsertChannel("data", "channel"),
			channelDetailService.VerifyJoinChannel("channel", "channelDetail"),
			messageService.CreateTextMessage("data"),
			middlewares.ResponseCreates("result"))
	}
}
