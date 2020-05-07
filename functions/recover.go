package functions

import (
	"fmt"
)

// UseRecover demonstrates use for recover function to handle panic mode
func UseRecover(){
	fmt.Println("Using Recover to handle panic calls")

	// Calling workFunction
	workFunction()

}

func workFunction(){
	// Defer function
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("defer function called in workFunction.")
			fmt.Println("Recovered from panic in workFunction")
		}
	}()

	// Calling panic Function
	panicFunction(0)
}

func panicFunction(i int){
	defer func(){
		// Defer functions are LIFO implementations
		fmt.Println("Defer function in panicFunction...", i)
	}()
	if i > 5 {
		fmt.Println("i > 5, calling panic function")
		panic(fmt.Sprintf("%v",i))
	}
	fmt.Printf("Processing %d in panic function\n",i)

	// Recursive call to panicFunction
	panicFunction(i + 1)
}