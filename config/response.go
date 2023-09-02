package config

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Result interface{} `json:"result"`
	Total  int64       `default:"1" json:"total"`
	Limit  int         `default:"1" json:"limit"`
	Offset int         `default:"0" json:"offset"`
}

type Response struct {
	Messages []Message    `json:"messages"`
	Data     ResponseData `json:"data"`
}

func NoData() ResponseData {
	return ResponseData{
		Result: nil,
		Total:  0,
		Limit:  0,
		Offset: 0,
	}
}

func (response *Response) GetSuccess(context *gin.Context) {
	context.JSON(200, response)
}

func (response *Response) Created(context *gin.Context) {
	context.JSON(201, response)
}

func (response *Response) BadRequest(context *gin.Context) {
	context.JSON(400, response)
}

func (response *Response) InternalServerError(context *gin.Context) {
	context.JSON(500, response)
}

func (response *Response) UnAuthorization(context *gin.Context) {
	context.JSON(401, response)
}

func (response *Response) UpdateSuccess(context *gin.Context) {
	context.JSON(202, response)
}
