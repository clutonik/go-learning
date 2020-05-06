package myerrors

import (
	"errors"
	"fmt"
	"log"
)

var errCustomRaw = errors.New("Value Error: I do not handle values greater than 5")
var errCustom = errors.New("Value Error: I do not handle values greater than 5")

// BasicError demonstrates use of errors package to create custom errors
func BasicError() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("My job was to recover from panic, but check the error message above...")
		}
	}()
	fmt.Println("Using errors package to write custom errors")

	_, err := customError(6)
	if err != nil {
		log.Println(err)
	}

	_, err = customErrorWithFormat(6)
	if err != nil {
		log.Panicln(err) // Create panic
	}
}

func customError(value int) (int, error) {
	limit := 5
	if value > limit {
		return 1, errCustomRaw
	}
	return 0, nil
}

func customErrorWithFormat(value int) (int, error) {
	limit := 5
	if value > limit {
		return 1, fmt.Errorf("%v, value provided: %v", errCustom, value)
	}
	return 0, nil
}
