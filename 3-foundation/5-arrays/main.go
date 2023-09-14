package main

import "fmt"

func main() {
	// Array in Go has fixed size
	var ageList [3]int

	ageList[0] = 27
	ageList[1] = 47
	ageList[2] = 25

	fmt.Printf("Age list: %v\n", ageList)
	fmt.Printf("Age list size: %v\n", len(ageList))

	for index, value := range ageList {
		fmt.Printf("Index: %v - Value: %v\n", index, value)
	}

}
