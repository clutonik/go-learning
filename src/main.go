package main

import (
	"exploreConstants"
	"exploreTypes"
	"fmt"
	"helloWorld"
	"zeroValue"
)

func main() {
	// Use helloWorld package to say Hello
	helloWorld.SayHello()
	fmt.Println("---------------------------")

	// Use zeroValue package to print default values for types
	zeroValue.ZeroValue()
	fmt.Println("---------------------------")

	// Explore type printing using fmt
	exploreTypes.PrintTypes()
	fmt.Println("---------------------------")

	// Explore constants
	exploreConstants.PrintConstants()
	fmt.Println("---------------------------")

}
