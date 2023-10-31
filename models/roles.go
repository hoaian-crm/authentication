package models

type Role struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`

	Permission []Permission `gorm:"many2many:permissions;" json:"permissions"`
}
