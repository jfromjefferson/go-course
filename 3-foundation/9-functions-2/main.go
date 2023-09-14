package main

import "fmt"

func main() {
	var result int = sum(1, 2, 3, 4)

	fmt.Println("Result:", result)

}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
