package user_dto

type UpdatePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password" binding:"min_length=8"`
}
