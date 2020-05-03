package main

import (
	"conditionals"
	"constants"
	"fmt"
	"functions"
	"goruntime"
	"helloworld"
	"interfaces"
	"iotas"
	"json"
	"loops"
	"maps"
	"slices"
	"structs"
	"tricks"
	"types"
	"zerovalue"
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
	functions.Callback()
	fmt.Println("---------------------------")

	// Explore Functions
	interfaces.Print()
	fmt.Println("---------------------------")

	// Exploring Tricks
	tricks.Print()
	fmt.Println("---------------------------")

	// Explore JSON
	json.Printjson()
	fmt.Println("---------------------------")

	// Go Runtime and routines
	//fmt.Println("Showing Race Condition...")
	//goruntime.Race()
	fmt.Println("Avoid Race Condition using Mutex...")
	goruntime.Mutex()
	fmt.Println("Avoid Race Condition using Atomic package...")
	goruntime.Atomic()
	fmt.Println("Showing Channels...")
	goruntime.Channel()
	goruntime.BufferedChannel()
	fmt.Println("Channel Types...")
	goruntime.ChannelTypes()
	fmt.Println("Using Directional Channels...")
	goruntime.DirectionalChannels()
	fmt.Println("---------------------------")
}
