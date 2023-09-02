package config

import (
	"fmt"
	"net/smtp"
)

type MailConfig struct {
	auth smtp.Auth
}

type MailData struct {
	To      string
	Subject string
	Content string
}

var MailSender MailConfig

func connect() {

}

func (mailSender *MailConfig) authentication() {
	auth := smtp.PlainAuth("", EnvirontmentVariables.MailUser, EnvirontmentVariables.MailPassword, EnvirontmentVariables.MailHost)
	mailSender.auth = auth
}

func (mailSender *MailConfig) SendMail(data MailData) {
	to := []string{data.To}
	msg := []byte("To: " + data.To + "\r\n" +
		"Subject: " + data.Subject + "\r\n" +
		"\r\n" +
		data.Content + "\r\n")

	smtp.SendMail(EnvirontmentVariables.MailHost+":"+EnvirontmentVariables.MailPort, mailSender.auth, EnvirontmentVariables.MailUser, to, msg)
}

func (mailSender *MailConfig) healthCheckMailService() {
	to := []string{"anhoai@playgroundvina.com"}
	msg := []byte("To: anhoai@playgroundvina.com\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", mailSender.auth, "hoaian412003@gmail.com", to, msg)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func InitializeMailSender() {
	MailSender = MailConfig{}
	MailSender.authentication()
	// MailSender.healthCheckMailService()
}
