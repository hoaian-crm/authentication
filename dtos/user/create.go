package user_dto

type Create struct {
	Email       string `gorm:"unique" binding:"email,must_unique=users" json:"email"`
	DisplayName string `json:"display_name" binding:"min_length=10"`
	Password    string `binding:"min_length=8" json:"password"`
	Avatar      string `json:"avatar"`
}
