package exploreLoops

import (
	"fmt"
)

func PrintLoops() {
	fmt.Println("Exploring Loops: ")

	x := 1
	// for loop
	fmt.Println("for loop with condition(equivalent to while loop in other languages): ")
	for x < 3 {
		fmt.Println("Iteration: ", x)
		x++
	}

	// For loop with init, condition, post statements
	fmt.Println("for loop with init, condition and post statements: ")
	for y := 1; y < 3; y++ {
		fmt.Println("Iteration: ", y)
	}

	// for without condition
	for {
		fmt.Println("for loop without condition.")	
        break
	}

}
