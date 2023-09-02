package user_dto

type ActiveUserDto struct {
	Email   string `json:"email" binding:"email"`
	OtpCode string `json:"otp_code"`
}
