package goruntime

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Atomic demonstrates race condition while using go routines
// Use go run -race command flag while testing this.
func Atomic() {
	fmt.Println("Avoid Race Condition using Atomic package...")
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	fmt.Println("Number of Go Routines: ", runtime.NumGoroutine())

	var counter int64 // Atomic works with int64, int32 etc.

	const routines = 100

	var waitGroup sync.WaitGroup
	waitGroup.Add(routines) // How many routines we want to add to waitgroup

	for i := 0; i < routines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) // Writing
			runtime.Gosched()            // Allowing processor to run other goroutines
			atomic.LoadInt64(&counter)   // Reading
			waitGroup.Done()
		}()
	}

	waitGroup.Wait() // Waiting for go routines to finish
	fmt.Println("Number of Go Routines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
