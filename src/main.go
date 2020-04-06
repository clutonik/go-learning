package main

import (
	conditionals "Conditionals"
	constants "Constants"
	functions "Functions"
	loops "Loops"
	maps "Maps"
	slices "Slices"
	structs "Structs"
	"fmt"
	helloworld "helloWorld"
	"interfaces"
	"iotas"
	"tricks"
	"types"
	zerovalue "zeroValue"
)

func main() {
	// Use helloWorld package to say Hello
	helloworld.Print()
	fmt.Println("---------------------------")

	// Use zeroValue package to print default values for types
	zerovalue.Print()
	fmt.Println("---------------------------")

	// Explore type printing using fmt
	types.Print()
	fmt.Println("---------------------------")

	// Explore constants
	constants.Print()
	fmt.Println("---------------------------")

	// Explore iota
	iotas.Print()
	fmt.Println("---------------------------")

	// Explore Conditionals
	conditionals.Print()
	fmt.Println("---------------------------")

	// Explore Loops
	loops.Print()
	fmt.Println("---------------------------")

	// Explore Slices
	slices.Print()
	fmt.Println("---------------------------")

	// Explore Maps
	maps.Print()
	fmt.Println("---------------------------")

	// Explore Structs
	structs.Print()
	fmt.Println("---------------------------")

	// Explore Functions
	functions.Print()
	fmt.Println("---------------------------")

	// Explore Functions
	interfaces.Print()
	fmt.Println("---------------------------")

	// Exploring Tricks
	tricks.Print()
	fmt.Println("---------------------------")
}
