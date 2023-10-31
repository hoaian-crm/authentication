package services

import (
	"main/config"
	"main/constants"
	"main/dtos"
	"main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PermissionService struct {
}

func (service PermissionService) Create(c *gin.Context) {

	db := c.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)
	data := c.MustGet("data").(models.Permission)

	db.Create(&data)

	response := config.Response{
		Data: config.ResponseData{
			Result: data,
		},
		Messages: []config.Message{config.Messages["create_success"]},
	}

	response.Created(c)
}

func (service PermissionService) List(c *gin.Context) {
	db := c.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)
	query := c.MustGet("query").(dtos.Query)

	if query.Keyword != "" {
		query.Keyword = "\"" + query.Keyword + "\""
		db.Where("policy like ? or name like ? or description like ?", query.Keyword, query.Keyword, query.Keyword)
	}

	var count int64
	db.Count(&count)

	db.Order(clause.OrderByColumn{Column: clause.Column{Name: query.Order}, Desc: query.Direction == "desc"})
	db.Limit(query.Limit)
	db.Offset(query.Offset)

	var result []models.Permission

	db.Find(&result)

	response := config.Response{
		Data: config.ResponseData{
			Result: result,
			Total:  count,
			Limit:  query.Limit,
			Offset: query.Offset,
		},
		Messages: []config.Message{
			config.Messages["get_success"],
		},
	}

	response.GetSuccess(c)

}
