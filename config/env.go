package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environtment struct {
	GormDSN       string
	DefaultAvatar string
	JwtSecret     string
	MailUser      string
	MailPassword  string
	MailHost      string
	MailPort      string
	RedisHost     string
	RedisPass     string
	RedisDb       string
}

var EnvirontmentVariables Environtment

func SetupEnvirontment() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	EnvirontmentVariables = Environtment{
		GormDSN:       os.Getenv("GORM_DSN"),
		DefaultAvatar: os.Getenv("DEFAULT_AVATAR"),
		JwtSecret:     os.Getenv("JWT_SECRECT"),
		MailUser:      os.Getenv("MAIL_USER"),
		MailPassword:  os.Getenv("MAIL_PASSWORD"),
		MailHost:      os.Getenv("MAIL_HOST"),
		MailPort:      os.Getenv("MAIL_PORT"),
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPass:     os.Getenv("REDIS_PASSWORD"),
		RedisDb:       os.Getenv("REDIS_DB"),
	}
}
