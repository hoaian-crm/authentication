package models

import "gorm.io/gorm"

type PointUse struct {
	gorm.Model
	/* UserID    int64        `json:"user_id"`
	User      User         `gorm:"foreginKey:UserId" json:"user"`
	UseCaseId int64        `json:"use_case_id"`
	UseCase   PointUseCase `gorm:"foreginKey:UseCaseId" json:"use_case"` */

	UserID    int64        `json:"user_id"`
	UseCaseId int64        `json:"use_case_id"`
	User      User         `gorm:"foreginKey:UserId" json:"user"`
	UseCase   PointUseCase `gorm:"foreginKey:UseCaseId" json:"use_case"`
}
