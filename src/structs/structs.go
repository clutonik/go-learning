package structs

import (
	"fmt"
)

type language struct {
	name    string
	country string
}

// Embed a struct to a new struct
type nativeLanguage struct {
	language
	native bool
}

// Print explores Structs in Go
func Print() {

	fmt.Println("Exploring Structs: ")

	// Use struct
	lang1 := language{
		name:    "Hindi",
		country: "India",
	}
	fmt.Println("Struct 1: ", lang1)

	// Access elements using '.' dot notation
	fmt.Println("Struct 1 using dot notation: ", "Language Name: ", lang1.name, ", Language Country: ", lang1.country)

	// Embedded Structs
	fmt.Println("Exploring Embedded Structs...")
	lang2 := nativeLanguage{
		language: lang1, //embedded type
		native:   true,
	}
	fmt.Println("Struct 2 using emebedded structs: ", lang2)
	fmt.Println("Embedded Struct using dot notation: ", "Language Name: ", lang2.name, ", Language Country: ", lang2.country, ", Native: ", lang2.native)

	// Another way of referencing a struct field which has been embedded into another struct
	fmt.Println("Another way of referencing values: <new-struct>.<parent-struct>.<field> : ", lang2.language.name)

	// Anonymous Structs : Structs which we need to use just once in our code
	person1 := struct {
		firstname string
		lastname  string
	}{
		firstname: "Steve",
		lastname:  "Jobs",
	}

	fmt.Println("Anonymous Struct: ", person1)

	/* Output:
	   Struct 1:  {Hindi India}
	   Struct 1 using dot notation:  Language Name:  Hindi , Language Country:  India
	   Exploring Embedded Structs...
	   Struct 2 using emebedded structs:  {{Hindi India} true}
	   Embedded Struct using dot notation:  Language Name:  Hindi , Language Country:  India , Native:  true
	   Another way of referencing values: <new-struct>.<parent-struct>.<field> :  Hindi
	   Anonymous Struct:  {Steve Jobs}
	*/
}
