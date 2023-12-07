package routes

import (
	"main/config"
	"main/constants"
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

		route.Use(middlewares.InitModel[models.Role](&models.Role{}, constants.DATABASE_META_KEY))

		route.POST("/", middlewares.BindBody[models.Role]("data"), roleService.Create)
		route.GET("/", middlewares.BindQuery[dtos.Query]("query"), roleService.List)
		route.PUT("/:roleId/permissions/:permissionId",
			middlewares.BindUri[role_dto.AttachRolePermission]("data"),
			middlewares.InitModel[models.Permission](&models.Permission{}, "permission_db"),
			roleService.AttachPermission)
		route.DELETE("/:roleId/permissions/:permissionId",
			middlewares.BindUri[role_dto.DetachPermission]("data"),
			middlewares.InitModel[models.Permission](&models.Permission{}, "permission_db"),
			roleService.DetachPermission)
		route.PATCH("/:roleId/permissions",
			middlewares.BindBody[role_dto.AttachRolePatchPermisison]("data"),
			middlewares.BindUri[role_dto.AttachRolePatchPermisison]("uri"),
			middlewares.InitModel[models.Permission](&models.Permission{}, "permission_db"),
			roleService.AttachPatchPermisson,
		)
		route.PUT("/:roleId", middlewares.BindUri[role_dto.UpdateRoleUri]("uri"), middlewares.BindBody[models.Role]("data"), roleService.Update)
		route.DELETE("/:roleId", middlewares.BindUri[role_dto.DeleteRoleUri]("uri"), roleService.Delete)
	}
}
