package services

import (
	"main/constants"
	messages_dto "main/dtos/messages"
	"main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MessageService struct {
}

func (MessageService) CreateTextMessage(key string) gin.HandlerFunc {
	return func(context *gin.Context) {
		user := context.MustGet("user").(models.User)
		data := context.MustGet("data").(messages_dto.Create)
		db := context.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)

		record := models.TextMessage{
			MessageAbstract: models.MessageAbstract{
				UserID:    int(user.ID),
				ChannelID: int(data.ChannelID),
				Type:      models.TEXT_MESSAGE,
			},
			Content: data.Content,
		}

		db.Create(&record)

		context.Set(constants.DATABASE_META_KEY, db)
		context.Set("result", record)
		context.Next()
	}
}
