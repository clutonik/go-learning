package helloWorld

import (
    "fmt"
)

func SayHello() {
    // Storing returned values (number of bytes written, errors)
    n, e := fmt.Println("Hello World") // Method definition at https://pkg.go.dev/fmt?tab=doc

    // Notice both Println statements below (Both are valid)
    fmt.Println("Number of bytes written: ", n); //with semi-colon (;)
    fmt.Println("Errors: ", e) // without semi-colon (;)
}
