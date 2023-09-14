package main

import "fmt"

func main() {
	var sumResult int = sum(5, 4)

	fmt.Println(sumResult)
	fmt.Println(sum(2, 4))
	fmt.Println(sum(9, 44))
	fmt.Println(magicFunction(12344568, 99874568))

}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func mod(a int, b int) int {
	return a % b
}

func magicFunction(a int, b int) (int, int, int, int, int) {
	var sumResult int = sum(a, b)
	var subResult int = sub(a, b)
	var mulResult int = mul(a, b)
	var divResult int = div(a, b)
	var modResult int = mod(a, b)

	return sumResult, subResult, mulResult, divResult, modResult
}
