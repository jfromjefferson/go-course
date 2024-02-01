package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "Test 1"
	}()

	msg := <-channel

	fmt.Println(msg)
}
