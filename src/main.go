package main

import (
	"github.com/clutonik/go-learning/exploreConstants"
	"github.com/clutonik/go-learning/exploreIota"
	"github.com/clutonik/go-learning/exploreTypes"
	"fmt"
	"github.com/clutonik/go-learning/helloWorld"
	"github.com/clutonik/go-learning/zeroValue"
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

}
