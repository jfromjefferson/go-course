package main

import "fmt"

func main() {
	value := 10

	var pointer *int = &value

	*pointer = 20

	fmt.Println("Pointer: ", *pointer)
	fmt.Println("value: ", value)

}
