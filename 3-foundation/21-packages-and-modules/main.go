package main

import (
	"mymodule/sum"
)

func main() {
	toSum := map[string]int{
		"John":  10,
		"Jane":  20,
		"James": 30,
		"Judy":  40,
	}

	resullt := sum.Sum(toSum)

	println(resullt)
}
