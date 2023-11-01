package routes

import (
	"main/config"
	"main/dtos"
	role_dto "main/dtos/role"
	"main/middlewares"
	"main/models"
	"main/services"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	route gin.RouterGroup
}

func NewRoleController(route *gin.RouterGroup) {

	config.Db.AutoMigrate(&models.Role{})

	roleService := services.RoleService{}

	route = route.Group("/roles")
	{

		route.Use(middlewares.InitModel[models.Role](&models.Role{}))
		route.Use(middlewares.Association("Permissions"))

		route.POST("/", middlewares.BindBody[models.Role]("data"), roleService.Create)
		route.GET("/", middlewares.BindQuery[dtos.Query]("query"), roleService.List)

		route.PUT("/:roleId/permissions/:permissionId", middlewares.BindUri[role_dto.AttachRolePermission]("data"), roleService.AttachPermission)
	}
}
