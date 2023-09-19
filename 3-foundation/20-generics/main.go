package main

func main() {
	m := map[string]int{
		"John":  10,
		"Jane":  20,
		"James": 30,
	}

	m2 := map[string]float64{
		"John":  10.5,
		"Jane":  20.5,
		"James": 30.5,
	}

	resullt := sum(m)
	resullt2 := sum(m2)

	println(resullt)
	println(resullt2)

	println(compare(10, 10))
	println(compare(10.5, 10.5))
	println(compare(10, 10.5))
	println(compare(10.5, 10))
}

func sum[T int | float64](m map[string]T) T {
	var total T

	for _, value := range m {
		total += value
	}

	return total
}

func compare[T comparable](a T, b T) bool {
	return a == b

}
