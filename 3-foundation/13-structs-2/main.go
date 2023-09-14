package main

import "fmt"

type Address struct {
	Street        string
	Number        int
	Neighbourhood string
	City          string
	Country       string
}

type Client struct {
	Name     string
	Age      int
	Address  Address
	isActive bool
	createAt string
}

func (client *Client) Deactivate() {
	client.isActive = false
}

func (client *Client) Activate() {
	client.isActive = true
}

func main() {
	address := Address{
		Street:        "P. Sherman",
		Number:        42,
		Neighbourhood: "Wallaby Way",
		City:          "Sydney",
		Country:       "Australia",
	}

	client := Client{
		Name:     "Jefferson",
		Age:      27,
		Address:  address,
		isActive: true,
		createAt: "2021-01-01 00:00:00",
	}

	fmt.Println(client)
	client.Deactivate()
	fmt.Println(client)

}
