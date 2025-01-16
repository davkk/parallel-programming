package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"rabbitmq-client/pkg/rabbitmq"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	username := flag.String("username", "", "select your username")
	flag.Parse()

	if username == nil || *username == "" {
		flag.Usage()
		os.Exit(1)
	}

	rmq := rabbitmq.Connect()
	defer rabbitmq.Close(rmq)

	err := rmq.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Panic("Failed to set QoS")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s> ", *username)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text != "" {
			err = rmq.Channel.Publish(
				"",             // exchange
				rmq.Queue.Name, // routing key
				false,          // mandatory
				false,          // immediate
				amqp.Publishing{
					ContentType:  "text/plain",
					DeliveryMode: amqp.Persistent,
					Body:         []byte(text),
					Headers: map[string]interface{}{
						"username": *username,
					},
				},
			)
			if err != nil {
				log.Panic("Failed to publish a message")
			}
		}
	}
}
