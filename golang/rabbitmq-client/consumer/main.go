package main

import (
	"log"

	"rabbitmq-client/pkg/rabbitmq"
)

func main() {
	rmq := rabbitmq.Connect()
	defer rabbitmq.Close(rmq)

	msgs, err := rmq.Channel.Consume(
		rmq.Queue.Name, // queue
		"",             // consumer
		false,          // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Panic("Failed to register a consumer")
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			username, ok := msg.Headers["username"].(string)
			if !ok {
				continue
			}
			log.Printf("%s> %s\n", username, msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
