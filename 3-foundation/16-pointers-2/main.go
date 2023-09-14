package main

import "fmt"

func main() {
	value := 10
	value1 := 20

	fmt.Println(sum(&value, &value1))
	fmt.Println(value)

}

func sum(a, b *int) int {
	*a = 50

	return *a + *b
}
