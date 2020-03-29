package exploreSlices

import (
	"fmt"
)

func PrintSlices() {
	// Composite Literals/Slices
	x := []int{1, 4, 6, 42, 56, 77}
	fmt.Println("Exploring Slices:: ")
	fmt.Printf("First Slice: %d\n", x)
	fmt.Println("Slice length: ", len(x))     // Method to get length of slice
	fmt.Println("Element at index 3: ", x[3]) // Way to access an element at a specific index

	fmt.Println("Looping over slice...")
	// Loop over slice elements
	for index, value := range x { // short declaration to declare two vars to access index and value of a slice
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Slicing a slice :)
	fmt.Println("Slicing the slice using colon(:) operator...")
	fmt.Println("First Slice: ", x[2:])  // Get rid of first 2 elements
	fmt.Println("Second Slice: ", x[:2]) // Print till index 2 (does not include element at index 2)
	fmt.Println("Third Slice: ", x[1:5]) // Prints elements from index 1 till index 5 (does not include element at index 5)

	// Creating a new slice by slicing a slice
	z := x[1:4]
	fmt.Println("Slice created by slicing a slice...")
	fmt.Println(z)

	// Adding elements to a slice
	fmt.Println("Appending elements to slice...")
	x = append(x, 2, 3, 10) // of form append(x,...y) (Variadic parameters)
	fmt.Println("Updated slice after adding elements: ", x)

	// Combining two slices together
	fmt.Println("Appending a slice to another slice...")
	y := []int{100, 200}
	x = append(x, y...) // append y... replaces y... with the elements in slice y
	fmt.Println("Updated slice after combining two slices: ", x)

	// Deleting elements from slice
	// Existing Slice x = [1 4 6 42 56 77 2 3 10 100 200]
	// Removing elements 6 and 42 from slice
	fmt.Printf("Removing %d and %d from slice %d\n", x[2], x[3], x)
	x = append(x[:2], x[4:]...) // Using syntax append(x,y...) to include all elements starting from position 4
	fmt.Println("Updated slice after deleting elements: ", x)

	// Create a slice using make (if we know the max length of slice, we will need)
	a := make([]int, 10, 100) // make is an inbuilt function. 10 is length and 100 is capacity.
	fmt.Println("Slice created using make: ", a)
	fmt.Println("Length of slice created using make: ", len(a))
	fmt.Println("Capacity of slice created using make: ", cap(a)) // If we reach the capacity of 100, runtime will create a new array with double capacity and will get rid of old array

	// Multi-dimensional Slice
	mds := [][]int{x, y}
	fmt.Println("Multi-dimensional slice: ", mds)

	// Looping over slice elements
	fmt.Println("Looping over slice elements...")
	for i, v := range mds {
		fmt.Printf("Index: %d, Value: %d", i, v)
	}
}
