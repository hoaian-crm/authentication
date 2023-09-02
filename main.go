package main

import (
	"fmt"
	"main/config"
	"main/routes"
	"main/validator"

	"github.com/gin-gonic/gin"
)

func main() {

	config.SetupEnvirontment()

	fmt.Printf("------------------- Connecting to database ----------------")

	config.ConnectDataBase()

	config.ConnectRedis()

	config.InitializeMailSender()

	r := gin.Default()

	validator.ValidatorBinding()

	api := r.Group("/api")
	{
		routes.NewUserController(api)
		routes.NewMailController(api)
		routes.NewPointUseCaseController(api)
		routes.NewPointUseController(api)
		routes.NewMessageController(api)
		routes.NewChannelController(api)
		routes.NewChannelDetailController(api)
	}

	r.Run()
}
