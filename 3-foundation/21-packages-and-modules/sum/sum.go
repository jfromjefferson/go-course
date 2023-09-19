package sum

func Sum[T int | float64](m map[string]T) T {
	var total T

	for _, value := range m {
		total += value
	}

	return total
}
