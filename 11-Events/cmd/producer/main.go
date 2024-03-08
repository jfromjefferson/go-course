package main

import (
	"fmt"
	"github.com/jfromjefferson/fcutils/pkg/rabbitmq"
	"os"
)

func main() {
	message := os.Args[1]
	fmt.Println("Producer start")
	channel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()

	err = rabbitmq.Publish(channel, message)
	if err != nil {
		panic(err)
	}
}
