package main

import "fmt"

func main() {
	priceList := []int{10}

	priceList = append(priceList, 11, 12, 131, 14, 15, 160, 17, 18, 19)

	fmt.Println(priceList)

}
