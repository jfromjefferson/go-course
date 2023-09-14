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

	client.isActive = false

	fmt.Println(client)
	fmt.Println(client.Address)

}
