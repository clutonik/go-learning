package goruntime

import (
	"fmt"
	"runtime"
	"sync"
)

// Race demonstrates race condition while using go routines
// Use go run -race command flag while testing this.
func Race() {
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	fmt.Println("Number of Go Routines: ", runtime.NumGoroutine())

	counter := 0

	const routines = 100

	var waitGroup sync.WaitGroup
	waitGroup.Add(routines) // How many routines we want to add to waitgroup

	for i := 0; i < routines; i++ {
		go func() {
			value := counter
			runtime.Gosched() //Allowing processor to run other goroutines
			value++
			counter = value
			waitGroup.Done()
		}()
	}

	waitGroup.Wait() // Waiting for go routines to finish
	fmt.Println("Number of Go Routines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
