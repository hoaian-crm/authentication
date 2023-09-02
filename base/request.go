package base

import (
	"main/config"

	"github.com/gin-gonic/gin"
)

func GetData[Dto any](context *gin.Context) Dto {
	var result Dto
	if err := context.ShouldBind(&result); err != nil {
		messages := config.MessagesBuilder(err)
		response := config.Response{
			Data:     config.NoData(),
			Messages: messages,
		}
		response.BadRequest(context)
	}
	return result
}
