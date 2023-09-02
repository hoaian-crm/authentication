package routes

import (
	"main/config"
	"main/models"

	"github.com/gin-gonic/gin"
)

func NewChannelController(route *gin.RouterGroup) {

	config.Db.AutoMigrate(&models.Channel{})

	route = route.Group("/channels")
	{

	}
}
