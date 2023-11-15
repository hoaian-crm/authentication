package services

import (
	"database/sql"
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
		db.Distinct()
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
	permissionDb := c.MustGet("permission_db").(*gorm.DB)
	data := c.MustGet("data").(role_dto.AttachRolePermission)

	role := models.Role{
		BaseModel: models.BaseModel{
			ID: data.RoleId,
		},
	}

	permission := models.Permission{
		BaseModel: models.BaseModel{
			ID: data.PermissionId,
		},
	}

	permissionDb.Where(&permission).Find(&permission)
	db.Model(&role).Association("Permissions").Append(&permission)
	db.Where(&role).Preload("Permissions").First(&role)

	response := config.Response{
		Data: config.ResponseData{
			Result: role,
		},
		Messages: []config.Message{
			config.Messages["update_success"],
		},
	}

	response.UpdateSuccess(c)

}

func (service RoleService) DetachPermission(c *gin.Context) {
	db := c.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)
	permissionDb := c.MustGet("permission_db").(*gorm.DB)
	data := c.MustGet("data").(role_dto.DetachPermission)

	role := models.Role{
		BaseModel: models.BaseModel{
			ID: data.RoleId,
		},
	}

	permission := models.Permission{
		BaseModel: models.BaseModel{
			ID: data.PermissionId,
		},
	}

	permissionDb.Where(&permission).Find(&permission)

	creator := config.Db
	creator.Model(&role).Association("Permissions").Delete(&permission)

	db.Where(&role).First(&role)

	response := config.Response{
		Data: config.NoData(),
		Messages: []config.Message{
			config.Messages["delete_success"],
		},
	}

	response.DeleteSuccess(c)
}

func (service RoleService) AttachPatchPermisson(c *gin.Context) {
	db := c.MustGet(constants.DATABASE_META_KEY).(*gorm.DB)
	permisisonDb := c.MustGet("permission_db").(*gorm.DB)
	data := c.MustGet("data").(role_dto.AttachRolePatchPermisison)
	uri := c.MustGet("uri").(role_dto.AttachRolePatchPermisison)

	data.RoleId = uri.RoleId
	role := models.Role{
		BaseModel: models.BaseModel{
			ID: data.RoleId,
		},
	}

	permissions := []models.Permission{}
	permisisonDb.Where("id in ?", data.PermissionIds).Find(&permissions)

	db.Model(&role).Association("Permissions").Append(&permissions)
	db.Where("id = ?", role.ID).Preload("Permissions").First(&role)

	response := config.Response{
		Data: config.ResponseData{
			Result: role,
		},
		Messages: []config.Message{
			config.Messages["update_success"],
		},
	}

	response.UpdateSuccess(c)

}
