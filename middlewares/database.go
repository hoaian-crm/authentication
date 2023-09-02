package middlewares

import (
	"main/config"
	"main/constants"
	"main/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitModel[Model any](instance *Model) gin.HandlerFunc {

	return func(context *gin.Context) {

		db, exits := context.Get(constants.DATABASE_META_KEY)
		if !exits {
			db = config.Db.Model(instance)
			context.Set(constants.DATABASE_META_KEY, db)
		}
		context.Next()
	}
}

func Creates(key string) gin.HandlerFunc {

	return func(context *gin.Context) {

		db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)
		data := context.MustGet(key)

		result := db.Create(utils.ToMap(data, "json"))
		context.Set(constants.DATABASE_META_KEY, result)
		context.Set(key, data)
		context.Next()
	}
}

func ResponseCreates(key string) gin.HandlerFunc {

	return func(context *gin.Context) {

		db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)
		record := context.MustGet(key)

		if db.Error != nil {
			message := config.Messages["create_failed"]
			message.Description = db.Error.Error()
			response := config.Response{
				Data:     config.NoData(),
				Messages: []config.Message{message},
			}
			response.BadRequest(context)
			return
		}

		response := config.Response{
			Data: config.ResponseData{
				Result: record,
			},
			Messages: []config.Message{config.Messages["create_success"]},
		}

		response.Created(context)
		context.Next()

	}
}

func Where[FilterDto any](key string) gin.HandlerFunc {

	return func(context *gin.Context) {

		db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)

		filter, exits := context.Get(key)

		if !exits {
			response := config.Response{
				Data:     config.NoData(),
				Messages: []config.Message{config.Messages["no_filter"]},
			}
			response.BadRequest(context)
			return
		}

		filter = filter.(FilterDto)

		result := db.Where(filter)
		context.Set(constants.DATABASE_META_KEY, result)
		context.Next()
	}
}

func ResponseFind[Model any]() gin.HandlerFunc {
	return func(context *gin.Context) {
		db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)

		if db.Error != nil {
			getFailedMessage := config.Messages["get_failed"]
			getFailedMessage.Description = db.Error.Error()
			response := config.Response{
				Data: config.ResponseData{
					Result: nil,
					Total:  db.RowsAffected,
					Limit:  0,
					Offset: 0,
				},
				Messages: []config.Message{
					getFailedMessage,
				},
			}

			response.BadRequest(context)
			return
		}

		getSuccessMessage := config.Messages["get_success"]
		var result []Model
		db.Find(&result)

		response := config.Response{
			Data: config.ResponseData{
				Result: result,
				Total:  db.RowsAffected,
				Limit:  0,
				Offset: 0,
			},
			Messages: []config.Message{
				getSuccessMessage,
			},
		}

		response.UpdateSuccess(context)
		context.Next()
	}
}

func Updates[UpdateDto any](key string) gin.HandlerFunc {
	return func(context *gin.Context) {

		db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)
		data, exits := context.Get(key)
		if !exits {
			response := config.Response{
				Data:     config.NoData(),
				Messages: []config.Message{config.Messages["no_data"]},
			}

			response.BadRequest(context)
		}

		result := db.Updates(data)

		context.Set(constants.DATABASE_META_KEY, result)
		context.Next()
	}
}

func ResponseUpdated() gin.HandlerFunc {
	return func(context *gin.Context) {

		db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)

		if db.Error != nil {
			updatedMessage := config.Messages["update_failed"]
			updatedMessage.Description = db.Error.Error()
			response := config.Response{
				Data: config.ResponseData{
					Result: nil,
					Total:  db.RowsAffected,
					Limit:  0,
					Offset: 0,
				},
				Messages: []config.Message{
					updatedMessage,
				},
			}

			response.UpdateSuccess(context)
			return
		}

		rowAffected, _ := utils.NumberToString(int(db.RowsAffected))
		updatedMessage := config.Messages["update_success"]

		updatedMessage.MetaData = map[string]string{
			"row_affected": rowAffected,
		}

		response := config.Response{
			Data: config.ResponseData{
				Result: nil,
				Total:  db.RowsAffected,
				Limit:  0,
				Offset: 0,
			},
			Messages: []config.Message{
				updatedMessage,
			},
		}

		response.UpdateSuccess(context)
		context.Next()
	}
}

func Deletes() gin.HandlerFunc {
	return func(context *gin.Context) {
		db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)
		result := db.Delete(map[string]interface{}{})

		context.Set(constants.DATABASE_META_KEY, result)
		context.Next()
	}
}

func ResponseDeleted() gin.HandlerFunc {
	return func(context *gin.Context) {

		db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)

		if db.Error != nil {
			updatedMessage := config.Messages["delete_failed"]
			updatedMessage.Description = db.Error.Error()
			response := config.Response{
				Data: config.ResponseData{
					Result: nil,
					Total:  db.RowsAffected,
					Limit:  0,
					Offset: 0,
				},
				Messages: []config.Message{
					updatedMessage,
				},
			}

			response.UpdateSuccess(context)
			return
		}

		rowAffected, _ := utils.NumberToString(int(db.RowsAffected))
		updatedMessage := config.Messages["delete_success"]

		updatedMessage.MetaData = map[string]string{
			"row_affected": rowAffected,
		}

		response := config.Response{
			Data: config.ResponseData{
				Result: nil,
				Total:  db.RowsAffected,
				Limit:  0,
				Offset: 0,
			},
			Messages: []config.Message{
				updatedMessage,
			},
		}

		response.UpdateSuccess(context)
		context.Next()
	}
}
