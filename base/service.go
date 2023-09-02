package base

import (
	"main/config"
	"main/dtos"

	"github.com/gin-gonic/gin"
)

type Service[Model any] struct {
	Repository Repository[Model]
}

func (service Service[Model]) CreateOne(context *gin.Context) {
	var dto Model
	if err := context.ShouldBind(&dto); err != nil {
		messages := config.MessagesBuilder(err)
		response := config.Response{
			Data:     config.NoData(),
			Messages: messages,
		}
		response.BadRequest(context)
		return
	}

	createdRecord := service.Repository.CreateOne(&dto)
	if createdRecord.Error != nil {
		response := config.Response{
			Data: config.NoData(),
			Messages: []config.Message{{
				Code:        -1,
				Description: createdRecord.Error.Error(),
			}},
		}
		response.InternalServerError(context)
		return
	}
	response := config.Response{
		Data: config.ResponseData{
			Limit:  1,
			Total:  1,
			Offset: 0,
			Result: dto,
		},
		Messages: []config.Message{},
	}

	response.Created(context)
	return
}

func (service Service[Model]) FindOne(context *gin.Context) {
	var filter Model
	if err := context.ShouldBind(&filter); err != nil {
		messages := config.MessagesBuilder(err)
		response := config.Response{
			Data:     config.NoData(),
			Messages: messages,
		}
		response.BadRequest(context)
		return
	}

	result, _ := service.Repository.FindOne(&filter)

	response := config.Response{
		Data: config.ResponseData{
			Limit:  1,
			Total:  1,
			Offset: 0,
			Result: result,
		},
		Messages: []config.Message{},
	}

	response.Created(context)
	return
}

func (service Service[Model]) FindAll(context *gin.Context) {
	query := GetData[dtos.Query](context)

	if query.Limit == 0 {
		query.Limit = 10
	}

	var filter Model

	result, total := service.Repository.FindAll(&filter, query)

	response := config.Response{
		Data: config.ResponseData{
			Result: result,
			Total:  total,
		},
		Messages: []config.Message{},
	}

	response.GetSuccess(context)
}
