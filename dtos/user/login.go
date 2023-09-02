package user_dto

type LoginDto struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"min_length=8"`
}
