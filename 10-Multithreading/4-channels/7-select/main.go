package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	select {
	case msg1 := <-c1:

		fmt.Println("Received", msg1)
	case msg2 := <-c2:
		fmt.Println("Received", msg2)
	}
}
