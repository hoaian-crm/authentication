package routes

import (
	"main/config"
	"main/constants"
	"main/dtos"
	"main/middlewares"
	"main/models"
	"main/services"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	route gin.RouterGroup
}

func NewPermissionController(route *gin.RouterGroup) {

	config.Db.AutoMigrate(&models.Permission{})

	permissionService := services.PermissionService{}

	route = route.Group("/permissions")
	{

		route.Use(middlewares.InitModel[models.Permission](&models.Permission{}, constants.DATABASE_META_KEY))

		route.POST("/", middlewares.BindBody[models.Permission]("data"), permissionService.Create)
		route.GET("/", middlewares.BindQuery[dtos.Query]("query"), permissionService.List)

	}
}
