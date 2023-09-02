package routes

import (
	"main/config"
	"main/models"

	"github.com/gin-gonic/gin"
)

func NewChannelDetailController(route *gin.RouterGroup) {

	config.Db.AutoMigrate(&models.ChannelDetail{})

	route = route.Group("/channel_details")
	{
	}
}
