package main

import (
	"fmt"
	"strings"
)

func main() {
	// salaries := make(map[string]int)
	salaries := map[string]int{
		"John":  125000,
		"Jane":  200000,
		"James": 300000,
	}

	// delete(salaries, "James")
	salaries["kassandra"] = 100000

	for key, value := range salaries {
		fmt.Printf("%s makes $%d\n", strings.Title(key), (value / 100))
	}

}
