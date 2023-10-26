package user_dto

import (
	"main/dtos"
)

type ListUserDto struct {
	dtos.Query
	DisplayName string `json:"displayName" form:"displayName"`
	Email       string `json:"email" form:"email"`
}

func (query *ListUserDto) SetDefaults() {
	query.Query.SetDefaults()

	if query.DisplayName == "" && query.Email {
		return
	}

	if query.DisplayName == "" {
		query.DisplayName = "/"
	}
	if query.Email == "" {
		query.Email = "/"
	}

}
