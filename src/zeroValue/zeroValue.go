package zerovalue

import (
	"fmt"
)

// Declare a variable called zerovar and of type int
// if we do not initialize a variable while defining type,
// GO assigns default zero value for that particular type
// Default zero values for Types:
// int = 0, floats = 0.0, strings = "", pointers, functions, interfaces, slices, channels, maps = nil
var zerovar int

// Print explores zeroValues in Go
func Print() {
	fmt.Println("Default Value: ", zerovar)
}
