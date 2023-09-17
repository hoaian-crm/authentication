package queue_base

import (
	"context"
	"encoding/json"
	"fmt"
	"main/config"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func New(queueName string) amqp091.Queue {
	q, _ := config.RabbitChannel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	return q
}

func Publish[T any](queueName string, data T) {
	dataEncoded, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error when convert data to send to json: %v\n", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := config.RabbitChannel.PublishWithContext(
		ctx,       // Context timeout 5 seconds
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(dataEncoded),
		}); err != nil {
		fmt.Printf("Error when publish message is: %v\n", err)
	}
}
