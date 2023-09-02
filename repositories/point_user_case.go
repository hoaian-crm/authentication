package repositories

import (
	"main/base"
	"main/models"
)

type PointUseCaseRepository struct {
	base.Repository[models.PointUseCase]
}
