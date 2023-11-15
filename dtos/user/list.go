package user_dto

import (
	"main/dtos"
)

type ListUserDto struct {
	dtos.Query
	Keyword string `json:"keyword" form:"keyword"`
}

func (query *ListUserDto) SetDefaults() {
	query.Query.SetDefaults()
}
