package routes

import (
	"main/config"
	pointuse_dto "main/dtos/point_use"
	"main/middlewares"
	"main/models"

	"github.com/gin-gonic/gin"
)

type PointUseController struct {
	route gin.RouterGroup
}

func NewPointUseController(route *gin.RouterGroup) {

	config.Db.AutoMigrate(&models.PointUse{})

	route = route.Group("/point_uses")
	{
		route.Use(middlewares.InitModel(&models.PointUse{}))

		route.POST("",
			middlewares.Authorization(),
			middlewares.BindBody[pointuse_dto.Create]("data"),
			middlewares.MergeBindData[models.User, pointuse_dto.Create]("user", "data", []string{"ID:UserID"}),
		)
	}
}
