package main

import "fmt"

type Client struct {
	name string
}

func (client *Client) walk() {
	client.name = "John Doe"

	fmt.Println(client.name, "is walking")
}

func main() {
	client := Client{name: "John"}

	client.walk()
	fmt.Println(client.name)

}
