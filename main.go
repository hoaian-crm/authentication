package main

import (
	"fmt"
	"main/config"
	"main/routes"
	"main/validator"
	docs "github.com/https://github.com/friendify-org/authentication-services/docs"

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
	}

	r.Run()
}
