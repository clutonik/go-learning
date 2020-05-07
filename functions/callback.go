package functions

import "fmt"

// Callback demonstrates use of callbacks in go
func Callback() {
	//Code to test callbacks - Which is passing function as an argument.
	fmt.Println("Exploring Callback:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total := sum(numbers...)
	fmt.Printf("\tTotal is %d\n", total)
	eventotal := evensum(sum, numbers...)
	fmt.Printf("\tTotal of even numbers is %d\n", eventotal)
}

func sum(si ...int) int {
	// This function accepts variadic parameters of type int
	total := 0
	for _, v := range si {
		total += v
	}

	return total
}

// Function accepting another function as argument
func evensum(f func(si ...int) int, values ...int) int {
	evens := []int{}
	for _, v := range values {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}

	return f(evens...)
}
