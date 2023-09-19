package main

func main() {
	var customerName interface{} = "John Doe"

	println(customerName.(string))
}
