package main

func main() {
	var x interface{} = 10
	var y interface{} = "Hello"

	showType(x)
	showType(y)

}

func showType(t interface{}) {
	switch t.(type) {
	case int:
		println("int")
	case string:
		println("string")
	default:
		println("unknown")
	}
}
