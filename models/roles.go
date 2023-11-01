package models

type Role struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`

	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}
