package main

import (
	"conditionals"
	"constants"
	"files"
	"fmt"
	"functions"
	"goruntime"
	"helloworld"
	"interfaces"
	"iotas"
	"json"
	"loops"
	"maps"
	"myerrors"
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
	functions.UseRecover()
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
	goruntime.Mutex()
	goruntime.Atomic()
	goruntime.Channel()
	goruntime.BufferedChannel()
	goruntime.ChannelTypes()
	goruntime.DirectionalChannels()
	goruntime.RangeAndCloseChannel()
	goruntime.SelectInChannel()
	goruntime.CommaOkIdiom()
	goruntime.UsingFanIn()
	goruntime.UsingFanOut()
	goruntime.ContextIntro()
	fmt.Println("---------------------------")

	// Working with files and demonstrating errors
	files.WriteToNewFile()

	// Using Custom Errors
	myerrors.BasicError()
}