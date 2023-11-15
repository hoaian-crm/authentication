package middlewares

import (
	"main/config"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func BindData[Dto any](context *gin.Context) Dto {
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

func BindBody[Dto any](key string) gin.HandlerFunc {
	return func(context *gin.Context) {
		var result Dto
		if err := context.ShouldBindJSON(&result); err != nil {
			var messages []config.Message
			if err.Error() == "EOF" {
				messages = []config.Message{config.Messages["invalid_body_data"]}
			} else {
				messages = config.MessagesBuilder(err)
			}
			response := config.Response{
				Data:     config.NoData(),
				Messages: messages,
			}
			response.BadRequest(context)
			context.Abort()
		}
		context.Set(key, result)
		context.Next()
	}
}

func BindQuery[Dto any](key string) gin.HandlerFunc {
	return func(context *gin.Context) {
		var result Dto
		if err := context.BindQuery(&result); err != nil {
			messages := config.MessagesBuilder(err)
			response := config.Response{
				Data:     config.NoData(),
				Messages: messages,
			}
			response.BadRequest(context)
		}
		context.Set(key, result)
		context.Next()
	}
}

func BindUri[Dto any](key string) gin.HandlerFunc {
	return func(context *gin.Context) {
		var result Dto
		if err := context.BindUri(&result); err != nil {
			messages := config.MessagesBuilder(err)
			response := config.Response{
				Data:     config.NoData(),
				Messages: messages,
			}
			response.BadRequest(context)
			context.Abort()
			return
		}
		context.Set(key, result)
		context.Next()
	}
}

func MergeBindData[DataSource any, DataTarget any](source string, target string, selects []string) gin.HandlerFunc {

	return func(context *gin.Context) {

		dataSource := context.MustGet(source).(DataSource)
		dataTarget := context.MustGet(target).(DataTarget)

		result := &dataTarget

		for _, key := range selects {
			data := strings.Split(key, ":")
			if len(data) < 2 {
				panic("Missing key to copy")
			}

			value := reflect.ValueOf(dataSource).FieldByName(data[0])

			reflect.ValueOf(result).Elem().FieldByName(data[1]).Set(value)
		}

		context.Set(target, *result)
	}
}
