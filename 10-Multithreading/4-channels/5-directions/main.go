package main

import "fmt"

func main() {
	hello := make(chan string)

	go publish("Hello", hello)
	reader(hello)
}

func publish(name string, hello chan<- string) {
	hello <- name
}

func reader(data <-chan string) {
	fmt.Println(<-data)
}
