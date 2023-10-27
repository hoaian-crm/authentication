package user_queue

import (
	"main/models"
	queue_base "main/queue"
)

func New() {
	queue_base.New("user_registered")
}

func UserRegister(data models.InternalUser) {
	queue_base.Publish("user_registered", data)
}
