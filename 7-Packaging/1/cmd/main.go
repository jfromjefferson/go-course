package main

import (
	"fmt"
	"gthub.com/jfromjefferson/go-course-7-1/math"
)

func main() {
	m := math.Math{A: 10, B: 20}

	fmt.Println("Result:", m.Add())
}
