package email_dto

type SendMailToUserDto struct {
	SendTo  string `json:"sendTo"`
	Content string `json:"content"`
	Subject string `json:"subject"`
}
