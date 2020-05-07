package types

import (
	"fmt"
)

// Variable at global scope
var z int = 4 // Recommended way of declaring a variable for global scope

// It is possible to create your own TYPES in GO.
// When we create CUSTOM TYPE, package name is added to its name i.e. package_name.custom_type
type id int

var userID id = 1 // user_id is a VARIABLE of CUSTOM TYPE exploreTypes.id and with value 1

// Print explores Types in Go.
func Print() {
	a := "I think I am a string" // Variable at local scope (using short declaration operator recommended for local vars)
	fmt.Println("%T:", a)        // This will not print the type
	fmt.Printf("%T: %v\n", a, a) // This will replace %T with TYPE of var named 'a'
	fmt.Printf("%T: %v\n", z, z) // This will replace %T with TYPE of var named 'z'

	// Custom TYPE
	fmt.Printf("%T\t%v\n", userID, userID)

}
