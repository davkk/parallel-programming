package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

func Connect() RabbitMQ {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panic("Failed to connect to RabbitMQ")
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Panic("Failed to open a channel")
	}

	queue, err := channel.QueueDeclare(
		"chat", // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Panic("Failed to declare a queue")
	}

	return RabbitMQ{conn, channel, &queue}
}

func Close(rmq RabbitMQ) {
	rmq.Conn.Close()
	rmq.Channel.Close()
}
