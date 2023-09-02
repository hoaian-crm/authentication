package services

import (
	"fmt"
	"main/config"
	messages_dto "main/dtos/messages"
	"main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChannelService struct {
}

func (ChannelService) UpsertChannel(dataKey string, outputKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		user := context.MustGet("user").(models.User)
		data := context.MustGet(dataKey).(messages_dto.Create)

		record := models.Channel{
			BaseModel: models.BaseModel{ID: uint(data.ChannelID)},
		}

		fmt.Printf("record: %v\n", record)

		if data.ChannelID == 0 {
			record = models.Channel{
				Type: models.BOT_CHAT,
				Name: data.Content,
			}
			config.Db.Create(&record)
			config.Db.Create(&models.ChannelDetail{
				ChannelID: int(record.ID),
				UserID:    int(user.ID),
			})
			data.ChannelID = int(record.ID)
			context.Set(dataKey, data)
			context.Set(outputKey, record)
		} else {
			result := config.Db.Model(&record).Where(&record).First(&record)
			if result.Error == gorm.ErrRecordNotFound {
				response := config.Response{
					Data: config.NoData(),
					Messages: []config.Message{
						config.Messages["not_found_channel"],
					},
				}
				response.BadRequest(context)
				context.Abort()
				return
			}
			context.Set(outputKey, record)
		}
		context.Next()
	}
}
