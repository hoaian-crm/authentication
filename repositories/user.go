package repositories

import (
	"main/base"
	"main/models"
)

type UserRepository struct {
	base.Repository[models.User]
}
