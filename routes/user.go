package routes

import (
	"main/base"
	"main/config"
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

		route.Use(middlewares.InitModel[models.User](&models.User{}))

		route.GET("/", middlewares.BindQuery[user_dto.ListUserDto]("query"), userService.ListUser)

		route.POST("/register", userService.Register)

		route.POST("/login", base.GetData[user_dto.LoginDto], userService.Login)

		route.GET("/profile", middlewares.Authorization(), userService.GetProfile)

		route.PUT("/active", middlewares.BindBody[user_dto.ActiveUserDto]("data"), userService.ActiveUser)

		route.PUT("/update_password", middlewares.Authorization(), userService.UpdatePassword)

	}
}
