package main

import (
	"fmt"
	"main/config"
	"main/docs"
	"main/routes"
	"main/validator"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

func main() {

	config.SetupEnvirontment()

	fmt.Printf("------------------- Connecting to database ----------------")

	config.ConnectDataBase()

	config.ConnectRedis()

	config.InitializeMailSender()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	validator.ValidatorBinding()

	api := r.Group("/api/v1")
	{
		routes.NewUserController(api)
		routes.NewMailController(api)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.PersistAuthorization(true)))

	r.Run()
}
