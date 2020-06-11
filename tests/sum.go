// Package tests demonstrates use of tests in Go
package tests

// Sum accepts variadic parameters of type int and returns their sum
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// Subtract accepts variadic parameters of type int and returns their substraction
func Subtract(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum -= v
	}
	return sum
}
