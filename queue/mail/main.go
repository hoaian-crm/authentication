package mail_queue

import (
	email_dto "main/dtos/email"
	queue_base "main/queue"
)

func New() {
	queue_base.New("mail")
}

func SendMailToUser(data email_dto.SendMailToUserDto) {
	queue_base.Publish("mail", data)
}
