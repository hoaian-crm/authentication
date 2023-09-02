package routes

import (
	"main/config"
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
		route.POST("/register", func(context *gin.Context) {
			userService.CreateOne(context)
		})
		route.POST("/login", func(context *gin.Context) {
			userService.Login(context)
		})
		route.GET("/profile", middlewares.Authorization(), func(context *gin.Context) {
			userService.GetProfile(context)
		})
		route.PUT("/active", func(context *gin.Context) {
			userService.ActiveUser(context)
		})
		route.PUT("/update_password", middlewares.Authorization(), func(context *gin.Context) {
			userService.UpdatePassword(context)
		})
	}
}
