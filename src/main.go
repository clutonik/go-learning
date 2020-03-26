package main

import (
	"exploreTypes"
	"helloWorld"
	"zeroValue"
)

func main() {
	// Use helloWorld package to say Hello
	helloWorld.SayHello()

	// Use zeroValue package to print default values for types
	zeroValue.ZeroValue()

	// Explore type printing using fmt
	exploreTypes.PrintTypes()

}
