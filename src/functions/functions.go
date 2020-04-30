package functions

import (
	"fmt"
)

// Print explores functions in go
func Print() {
	fmt.Println("Exploring Functions...")
	basicFunction()                    // basic function with a print statement
	defer basicFunction()              // basic function again with defer to execute it at the end
	pmFunction("Test Parameter")       // parameterized function
	rv := returnFunction("Steve Jobs") // Calling function which returns a string
	fmt.Println(rv)
	rv2, rv3 := returnFunction2("Bill Gates") // Calling function which returns 2 strings
	fmt.Println(rv2, rv3)

	// Using variadic function (ability to pass zero to unlimited numbers)
	sum := variadicFunction(1, 2, 3, 4, 5, 6)
	fmt.Printf("\tSum using variadic function(passing numbers): %d\n", sum)

	// Passing slice to variadic function instead of passing numbers
	xi := []int{1, 2, 3, 4, 5, 6}
	sum = variadicFunction(xi...)
	fmt.Printf("\tSum using variadic function(passing slice): %d\n", sum)

	// Anonymous Function
	func(x int) {
		fmt.Printf("\tPassed in value to anonymous function is: %d\n", x)
	}(10)

	// Func expressions
	fe := func(name string) {
		fmt.Printf("\tName passed in to func expression: %s\n", name)
	}

	// Calling the function declared earlier
	fe("Clutonik")
}

// Function Signature
// func (r receiver) identifier (parameter(s)) (return(s)) { code }
// Function with a print statement
func basicFunction() {
	fmt.Println("\tBasic function without any parameters, return values")
}

// function with parameters
func pmFunction(p string) {
	fmt.Println("\tParameterized function with parameter: ", p)
}

// Function with parameter and return value
func returnFunction(p string) string {
	return fmt.Sprint("\tParameter from function returnFunction: ", p)
}

// Function with parameter and multiple return values
func returnFunction2(p string) (string, string) {
	a := fmt.Sprint("\tParameter from function returnFunction2: ", p)
	b := "is real"
	return a, b
}

func variadicFunction(numbers ...int) int {
	sum := 0
	for v := range numbers {
		sum += v
	}
	return sum
}

/* Output:
Exploring Functions...
    Basic function without any parameters, return values
    Parameterized function with parameter:  Test Parameter
    Parameter from function returnFunction: Steve Jobs
	Parameter from function returnFunction2: Bill Gates is real
	Sum using variadic function(passing numbers): 15
	Sum using variadic function(passing slice): 15
	Basic function without any parameters, return values
*/
