package main

import (
	"exploreConditionals"
	"exploreConstants"
	"exploreFunctions"
	"exploreIota"
	"exploreLoops"
	"exploreMaps"
	"exploreSlices"
	"exploreStructs"
	"exploreTypes"
	"exploremethods"
	"fmt"
	"helloWorld"
	"tricks"
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

	// Explore iota
	exploreIota.PrintIota()
	fmt.Println("---------------------------")

	// Explore Conditionals
	exploreConditionals.PrintConditionals()
	fmt.Println("---------------------------")

	// Explore Loops
	exploreLoops.PrintLoops()
	fmt.Println("---------------------------")

	// Explore Slices
	exploreSlices.PrintSlices()
	fmt.Println("---------------------------")

	// Explore Maps
	exploreMaps.PrintMaps()
	fmt.Println("---------------------------")

	// Explore Maps
	exploreStructs.PrintStructs()
	fmt.Println("---------------------------")

	// Explore Functions
	exploreFunctions.PrintFunctions()
	fmt.Println("---------------------------")

	// Explore Functions
	exploremethods.PrintMethods()
	fmt.Println("---------------------------")

	// Exploring Tricks
	tricks.PrintTricks()
	fmt.Println("---------------------------")
}
