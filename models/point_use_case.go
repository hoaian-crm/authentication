package models

import "gorm.io/gorm"

type PointUseCase struct {
	gorm.Model
	Value       int64  `json:"value" binding:"is_number"`
	Description string `json:"description" binding:"is_not_empty"`
}
