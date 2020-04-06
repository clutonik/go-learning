package maps

import (
	"fmt"
)

// Print explores maps in Go
func Print() {
	fmt.Println("Exploring Maps: ")

	// Declaring Map
	x := map[string]string{
		"Name":     "Clutonik",
		"Business": "software",
	}

	fmt.Println("Printing map: ", x) // Print map
	fmt.Println("Accessing Keys and Values...")
	fmt.Println("Key: ", x["Name"])
	fmt.Println("Value: ", x["Business"])

	// We will get zeroValue for a key which does not exist
	fmt.Println("Accessing a key which does not exist...")
	fmt.Println("Registration: ", x["Registration"]) // this will print ""

	// Storing access element result in a variable
	fmt.Println("Storing access element request result in variables...")
	value, status := x["Registration"] // Accessing a value of a key in map returns two vars, value and state (true/false)
	fmt.Println("value=", value, "status=", status)

	// If condition to check if a key exists
	fmt.Println("If condition to check if a key exists...")
	if v, ok := x["Name"]; ok { // This is also known as "comma ok" idiom in GO
		fmt.Println("OK: ", ok)
		fmt.Println("Value: ", v)
	}

	// Adding an element to the map
	x["Registration"] = "done"

	// Loop over keys and values in a map
	fmt.Println("Looping over map keys...")
	for key, value := range x {
		fmt.Println("Key: ", key, ", Value: ", value)
	}

	// Deleting a key from a map
	fmt.Printf("Deleting '%s' from %s\n", x["Registration"], x)
	delete(x, "Registration")
	fmt.Println("Updated map: ", x)

	// Deleting a key which does not exist
	fmt.Printf("Trying to Delete ['Registration'] from %s which does not exist\n", x)

	if _, ok := x["Registration"]; ok {
		delete(x, "Registration")
		fmt.Println("Updated map: ", x)
	} else {
		fmt.Println("No deletion happened, as the key does not exists in map...")
	}
}
