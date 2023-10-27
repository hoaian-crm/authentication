package main

import (
	"main/config"
	mail_queue "main/queue/mail"
	"main/routes"
	"main/validator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

func main() {

	config.SetupEnvirontment()

	config.ConnectDataBase()

	config.ConnectRedis()

	config.InitializeMailSender()

	config.ConnectQueue()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
	}))

	validator.ValidatorBinding()

	api := r.Group("/api/v1")
	{
		routes.NewUserController(api)
	}

	mail_queue.New()

	r.Run()
}
