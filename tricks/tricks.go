package tricks

import (
	"fmt"
)

// Declaring two variables
var x, y int

// Print explores tricks I learned in Go
func Print() {
	fmt.Println("Code samples demonstrating few tricks: ")
	x, y = 42, 44 // Assigning values to two variables at the same time

	// Tricks:
	// Trick-1: Swapping values of two variables
	fmt.Println("Trick-1: Code to swap values of two variables...")
	fmt.Println("\tValues before swapping: ", x, y)
	x, y = y, x // Swapping values
	fmt.Println("\tValues after swapping: ", x, y)
}
