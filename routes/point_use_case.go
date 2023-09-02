package routes

import (
	"main/config"
	"main/dtos/point_use_case"
	"main/middlewares"
	"main/models"
	"main/services"

	"github.com/gin-gonic/gin"
)

type PointUseCaseController struct {
	route gin.RouterGroup
}

func NewPointUseCaseController(route *gin.RouterGroup) {

	config.Db.AutoMigrate(&models.PointUseCase{})

	pointUseCaseService := services.PointUseCaseService{}

	route = route.Group("/point_use_cases")
	{
		route.Use(middlewares.InitModel[models.PointUseCase](&models.PointUseCase{}))

		route.POST("",
			middlewares.Authorization(),
			middlewares.BindBody[pointusecase_dto.Create]("data"),
			middlewares.Creates("data"),
			middlewares.ResponseCreates("data"),
		)
		route.GET("", middlewares.Authorization(), func(context *gin.Context) {
			pointUseCaseService.FindAll(context)
		})
		route.GET(":id",
			middlewares.Authorization(),
			middlewares.BindUri[pointusecase_dto.Uri]("uri"),
			middlewares.Where[pointusecase_dto.Uri]("uri"),
			middlewares.ResponseFind[models.PointUseCase](),
		)
		route.PUT(":id",
			middlewares.Authorization(),
			middlewares.BindUri[pointusecase_dto.Uri]("uri"),
			middlewares.BindBody[pointusecase_dto.Data]("data"),
			middlewares.Where[pointusecase_dto.Uri]("uri"),
			middlewares.Updates[pointusecase_dto.Data]("data"),
			middlewares.ResponseUpdated(),
		)
		route.DELETE(":id",
			middlewares.Authorization(),
			middlewares.BindUri[pointusecase_dto.Uri]("uri"),
			middlewares.Where[pointusecase_dto.Uri]("uri"),
			middlewares.Deletes(),
			middlewares.ResponseDeleted(),
		)
	}
}
