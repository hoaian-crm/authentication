package routes

import (
	"main/config"
	"main/constants"
	user_dto "main/dtos/user"
	"main/middlewares"
	"main/models"
	"main/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	route gin.RouterGroup
}

func NewUserController(route *gin.RouterGroup) {

	config.Db.AutoMigrate(&models.User{})

	userService := services.UserService{}

	route = route.Group("/users")
	{

		route.Use(middlewares.InitModel[models.User](&models.User{}, constants.DATABASE_META_KEY))
		route.Use(middlewares.Association("Roles"))

		route.GET("/list", middlewares.BindQuery[user_dto.ListUserDto]("query"), userService.ListUser)

		route.POST("/register", middlewares.BindBody[models.User]("data"), userService.Register)

		route.POST("/login", middlewares.BindBody[user_dto.LoginDto]("data"), userService.Login)

		route.GET("/profile", middlewares.Authorization(), userService.GetProfile)

		route.PUT("/active", middlewares.BindBody[user_dto.ActiveUserDto]("data"), userService.ActiveUser)

		route.PUT("/update_password", middlewares.Authorization(), userService.UpdatePassword)

	}
}
