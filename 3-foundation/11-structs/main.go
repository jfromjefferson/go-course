package main

import "fmt"

type Client struct {
	Name     string
	Age      int
	Address  string
	isActive bool
	createAt string
}

func main() {
	client := Client{
		Name:     "Jefferson",
		Age:      27,
		Address:  "P. Sherman, 42 Wallaby Way, Sydney",
		isActive: true,
		createAt: "2021-01-01 00:00:00",
	}

	client.isActive = false

	fmt.Println(client)

}
