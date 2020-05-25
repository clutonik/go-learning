package tests

// Sum accepts variadic parameters of type int and returns their sum
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
