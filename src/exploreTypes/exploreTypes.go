package exploreTypes

import (
    "fmt"
)
// Variable at global scope
var z int = 4 // Recommended way of declaring a variable for global scope

func PrintTypes() {
    a := "I think I am a string" // Variable at local scope (using short declaration operator recommended for local vars)
    fmt.Println("%T:",a) // This will not print the type 
    fmt.Printf("%T: %v\n",a,a) // This will replace %T with TYPE of var named 'a'
    fmt.Printf("%T: %v\n",z,z) // This will replace %T with TYPE of var named 'z'
}
