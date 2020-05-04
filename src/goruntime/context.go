package goruntime

import (
	"context"
	"fmt"
	"time"
)

// ContextIntro demonstrates basic use of context in go routines
func ContextIntro() {
	fmt.Println("Using most basic version of context...")
	// Initializing subcontext by passing parent context which is context.Background()
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Coming out of go routine as context is cancelled.")
				return
			default:
				// If context is not cancelled increment n
				n++
				fmt.Println("n incremented: ", n)
				time.Sleep(time.Millisecond * 300) // Relative to how much do we wait in main go routine
			}
		}
	}()

	// Adding sleep in main goroutine so that go routine using ctx
	// could run for a while
	time.Sleep(time.Second * 3) // This will allow go routine above to run until n=10
	cancel()                    // Cancelling context which will cause go routine created above to exit

	fmt.Println(ctx.Err())
	// Adding some extra wait to let other go routines finish if there is any running
	time.Sleep(time.Second * 2)
}
