package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deletedAt" json:"deletedAt"`
}
