package middlewares

import (
	"main/config"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {

	return func(context *gin.Context) {
		// authorizationHeaders := strings.Split(context.GetHeader("Authorization"), " ")
		// if len(authorizationHeaders) < 2 {
		// 	response := config.Response{
		// 		Data:     config.NoData(),
		// 		Messages: []config.Message{config.Messages["missing_token"]},
		// 	}
		// 	response.UnAuthorization(context)
		// 	context.Abort()
		// 	return
		// }
		// token := authorizationHeaders[1]
		// userRecord := models.User{}
		// err := userRecord.VerifyToken(token)
		// if err != nil {
		// 	response := config.Response{
		// 		Data:     config.NoData(),
		// 		Messages: []config.Message{config.Messages["invalid_access_token"]},
		// 	}
		// 	response.UnAuthorization(context)
		// 	context.Abort()
		// 	return
		// }
		// context.Set("user", userRecord)
		userIdRaw := context.GetHeader("userId")
		value, err := utils.StringToNumber(userIdRaw)
		if err != nil {
			response := config.Response{
				Data:     config.NoData(),
				Messages: []config.Message{config.Messages["missing_token"]},
			}
			response.UnAuthorization(context)
			context.Abort()
			return
		}
		context.Set("userId", int64(value))
		context.Next()
	}
}
