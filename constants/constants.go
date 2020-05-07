package constants

import (
	"fmt"
)

// Declaring constants by keeping ease of programming in mind
const (
	a int    = 24   // specifying type
	b        = 34.2 // giving compiler flexibility to figure out type
	c string = "Constant in go"
)

// Declaring constants individually
const d = 32
const e = 34.4
const f = "Both ways work"

// Print explores constants in Go
func Print() {
	fmt.Println("Constant from first set: ", a)
	fmt.Println("Constant from second set: ", d)
}
