package config

import (
	"github.com/rabbitmq/amqp091-go"
)

var RabbitChannel *amqp091.Channel

func ConnectQueue() {
	conn, err := amqp091.Dial(EnvirontmentVariables.RabbitMqUri)
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	RabbitChannel = channel
	DeclareRequiredQueue()
}

func DeclareRequiredQueue() {

	for _, channelName := range EnvirontmentVariables.ChannelRequires {
		RabbitChannel.QueueDeclare(
			channelName, // name
			false,       // durable
			false,       // delete when unused
			false,       // exclusive
			false,       // no-wait
			nil,         // arguments
		)
	}
}
