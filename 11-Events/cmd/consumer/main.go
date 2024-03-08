package main

import (
	"fmt"
	"github.com/jfromjefferson/fcutils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Consumer start")
	channel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()

	messages := make(chan amqp.Delivery)
	go rabbitmq.Consume(channel, messages)
	for message := range messages {
		fmt.Println(string(message.Body))
		message.Ack(false)
	}
}
