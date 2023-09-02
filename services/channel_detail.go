package services

import (
	"main/config"
	"main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChannelDetailService struct {
}

func (ChannelDetailService) VerifyJoinChannel(dataKey string, outputKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		user := context.MustGet("user").(models.User)
		data := context.MustGet(dataKey).(models.Channel)

		query := models.ChannelDetail{
			ChannelID: int(data.ID),
			UserID:    int(user.ID),
		}

		channelDetail := models.ChannelDetail{}

		result := config.Db.Model(&query).Where(&query).First(&channelDetail)
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

		context.Set(outputKey, channelDetail)

		context.Next()
	}
}
