package services

import (
	"database/sql"
	"fmt"
	"main/config"
	"main/constants"
	"main/dtos"
	role_dto "main/dtos/role"
	"main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoleService struct {
}

func (service RoleService) Create(c *gin.Context) {

	db := c.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)
	data := c.MustGet("data").(models.Role)

	db.Create(&data)

	response := config.Response{
		Data: config.ResponseData{
			Result: data,
		},
		Messages: []config.Message{config.Messages["create_success"]},
	}

	response.Created(c)
}

func (service RoleService) List(c *gin.Context) {
	db := c.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)
	query := c.MustGet("query").(dtos.Query)

	query.SetDefaults()

	if query.Keyword != "" {
		query.Keyword = "%" + query.Keyword + "%"
		db.Joins("left join role_permissions on role_permissions.role_id = roles.id")
		db.Joins("left join permissions on role_permissions.permission_id = permissions.id")
		db.Where("roles.name like @Keyword or roles.description like @Keyword or permissions.name like @Keyword or permissions.description like @Keyword or permissions.policy like @Keyword",
			sql.Named("Keyword", query.Keyword))
	}

	var count int64
	db.Count(&count)

	if query.Order != "" {
		db.Order(clause.OrderByColumn{Column: clause.Column{Name: query.Order}, Desc: query.Direction == "desc"})
	}
	db.Limit(query.Limit)
	db.Offset(query.Offset)

	result := []models.Role{}

	db.Preload("Permissions").Find(&result)

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

func (service RoleService) AttachPermission(c *gin.Context) {
	db := c.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)
	data := c.MustGet("data").(role_dto.AttachRolePermission)

	role := models.Role{
		BaseModel: models.BaseModel{
			ID: data.RoleId,
		},
	}

	db.Where(&role).First(&role)

	fmt.Printf("role: %v\n", role)
}
