package email_dto

type SendMailToUserDto struct {
	UserId  int    `json:"userId"`
	Content string `json:"content"`
	Subject string `json:"subjet"`
}
