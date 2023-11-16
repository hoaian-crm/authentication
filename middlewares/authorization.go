package middlewares

import (
	"main/utils"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {

	return func(context *gin.Context) {
		userIdRaw := context.GetHeader("UserId")
		value, err := utils.StringToNumber(userIdRaw)
		if err != nil {
			value = 1
		}
		context.Set("UserId", value)
		context.Next()
	}
}
