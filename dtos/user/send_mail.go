package user_dto

type SendVerifyMail struct {
	Email string `json:"email" binding:"email"`
}
