package myerrors

import (
	"errors"
	"fmt"
	"log"
)

// Error messages as vars
var errCustom = errors.New("Value Error: I do not handle values greater than ")

// Error messages as types (A bit advanced)
// Creating struct to create a custom error type for log files
// Implement Error() method which returns a string to implement error interface
// https://golang.org/pkg/builtin/#error
type errLog struct {
	err            error
	logLevel       string
	originFunction string
}

// Implementing error interface by attaching Error() to errLog struct created above.
func (e errLog) Error() string {
	return fmt.Sprintf("%v - %v - %v", e.logLevel, e.originFunction, e.err)
}

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

	// Using formated error message
	_, err = customErrorWithFormat(6)
	if err != nil {
		log.Println(err) // Create panic
	}

	// Using custom type (using struct)
	_, err = customErrorForLogs(6)
	if err != nil {
		log.Panicln(err) // Create panic
	}
}

func customError(value int) (int, error) {
	limit := 5
	if value > limit {
		return 1, errCustom
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

func customErrorForLogs(value int) (int, error) {
	limit := 5
	if value > limit {
		formatedMessage := fmt.Errorf("%v, value provided: %v", errCustom, value)
		return 1, errLog{
			err:            formatedMessage,
			logLevel:       "ERROR",
			originFunction: "customErrorForLogs",
		}
	}
	return 0, nil
}
