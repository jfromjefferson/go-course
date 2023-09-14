package main

import "fmt"

func main() {
	var result int = sum(1, 2, 3, 4)

	resultTemp := func() int {
		return sum(1, 2, 3, 4) * 2
	}()

	fmt.Println("Result:", result)
	fmt.Println("Result temp:", resultTemp)
	fmt.Println("Total:", result+resultTemp)

}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
