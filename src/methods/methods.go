package methods

import (
	"fmt"
)

type language struct {
	name    string
	country string
}

// method
func (l language) speak() {
	fmt.Printf("\t%s is spoken in %s\n", l.name, l.country) // Accessing fields of type language
}

// Print demonstrates use of methods
func Print() {

	fmt.Println("Exploring Methods: ")
	lang1 := language{
		name:    "Hindi",
		country: "India",
	}

	lang1.speak() // Usage is similar to instance methods in other languages
}

/* Output:
   Exploring Methods:
   		Hindi is spoken in India
*/
