package json

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname string // These fields have to be in Inital Case
	Lastname  string
	Address   string
}

type convertedPerson struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Address   string `json:"Address"`
}

// Printjson demonstrates json library
func Printjson() {
	person1 := person{
		Firstname: "Narendra",
		Lastname:  "Modi",
		Address:   "India",
	}

	person2 := person{
		Firstname: "Arvind",
		Lastname:  "Kejriwal",
		Address:   "India",
	}

	people := []person{
		person1,
		person2,
	}

	fmt.Println("People: ", people)

	// Using json standard library to Create slice of bytes []byte (json)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(bs)) // Will return slice of bytes if Type is not casted

	var convertedPeople []convertedPerson

	// Unmarshal JSON which points to a datastructure (pointer)
	uerr := json.Unmarshal(bs, &convertedPeople) // Accepts slice of bytes
	if uerr != nil {
		fmt.Println("Error: ", uerr)
	}

	fmt.Println("Converted data: ", convertedPeople)

}
