package goruntime

import (
	"fmt"
	"sync"
	"time"
)

// Channel demonstrates use of channels in go routines
func Channel() {
	fmt.Println("Showing Channels...")
	channel := make(chan int)

	go func() {
		channel <- 40
	}()

	fmt.Println("Value in Channel: ", <-channel)
}

// BufferedChannel demonstrates use of buffered channels
func BufferedChannel() {
	channel := make(chan int, 1)

	channel <- 41

	fmt.Println("Value from buffered channel: ", <-channel)
}

// ChannelTypes demonstrates use of different channel types
func ChannelTypes() {
	fmt.Println("Channel Types...")
	readChannel := make(<-chan int)
	writeChannel := make(chan<- int)

	fmt.Printf("Read Channel: %T\n", readChannel)
	fmt.Printf("Write Channel: %T\n", writeChannel)
}

// DirectionalChannels demonstrates use of directional channels in Go
func DirectionalChannels() {
	fmt.Println("Using Directional Channels...")
	channel := make(chan int)

	// send
	go sendChannel(channel)

	// Receive
	receiveChannel(channel)

}

func sendChannel(channel chan<- int) {
	fmt.Println("Sending int value to channel")
	channel <- 42
}

func receiveChannel(channel <-chan int) {
	fmt.Println("Value in Channel: ", <-channel)
}

// RangeAndCloseChannel demonstrates ranging over channel until it is closed.
func RangeAndCloseChannel() {
	fmt.Println("Ranging over channel to read values...")
	channel := make(chan int)

	// Send values to channel
	go func() {
		for i := 0; i < 5; i++ {
			channel <- i
		}
		close(channel)
	}()

	// Read values from Channel until it is closed
	for v := range channel {
		fmt.Println("Value from Channel: ", v)
	}
}

// SelectInChannel Demonstrates use of select statement while using channels
func SelectInChannel() {
	fmt.Println("Using Select Statement to read values from channels...")
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	quitChannel := make(chan int)

	// send values to channel
	go sendUsingMultipleChannels(evenChannel, oddChannel, quitChannel)

	// read Values from channel
	receiveUsingMultipleChannels(evenChannel, oddChannel, quitChannel)

	fmt.Println("end of SelectInChannel function")

}

func receiveUsingMultipleChannels(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Value from Even Channel: ", v)
		case v := <-o:
			fmt.Println("Value from Odd Channel: ", v)
		case v := <-q:
			fmt.Println("Value from Quit Channel: ", v)
			return
		}
	}
}

func sendUsingMultipleChannels(e, o, q chan<- int) {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e) TODO: Read why close causes multiple 0 values to be printed
	// close(o)
	q <- 0
}

// CommaOkIdiom demonstrates use of comma ok idiom to check if
// a channel is closed
func CommaOkIdiom() {
	fmt.Println("Using Comma Ok Idiom to check channel status and value...")
	channel := make(chan int)

	// Send value to channel
	go func() {
		channel <- 42
		close(channel)
	}()

	value, ok := <-channel
	fmt.Printf("Value: %d, Channel ok ? : %t\n", value, ok)

	// Checking value and status again after the value has been already read
	value, ok = <-channel
	fmt.Printf("Value: %d, Channel ok ? : %t\n", value, ok)
}

// UsingFanIn demonstrates using a dedicated channel to store
// results
func UsingFanIn() {
	fmt.Println("Using a FanIn Channel to store results...")

	// Initialize three channels here
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	fanInChannel := make(chan int) // Channel to store results

	// Calling method to send values to even and odd channels
	go sendValuesToChannels(evenChannel, oddChannel)

	// Calling method to read values from two channels and storing
	// those values in fanInChannel
	go receiveValuesIntoFanInChannel(evenChannel, oddChannel, fanInChannel)

	// Ranging over fanInChannel to read values
	for value := range fanInChannel {
		fmt.Println("Value from fanInChannel: ", value)
	}
}

func sendValuesToChannels(e, o chan<- int) {
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func receiveValuesIntoFanInChannel(e, o <-chan int, f chan<- int) {
	// Using waitgroups as we need to execute go routines which should
	// be completed before we exit this method
	var wg sync.WaitGroup
	wg.Add(2) // Adding to routines which we need to wait for

	// Ranging over even channel
	go func() {
		for v := range e {
			f <- v // Storing even values to fanIn channel
		}
		wg.Done()
	}()

	// Ranging over odd channel
	go func() {
		for v := range o {
			f <- v // Storing Odd Values to fanIn Channel
		}
		wg.Done()
	}()

	wg.Wait()
	close(f)
}

// UsingFanOut demonstrates use of FanOut process to start
// go routines for time consuming work by reading values from
// another channel
func UsingFanOut() {
	fmt.Println("Using Fan out channels...")

	// Create two channels
	sendChannel := make(chan int)
	fanOutChannel := make(chan int)

	// Send values to sendChannel
	go populate(sendChannel)

	// Fanning out work to timeConsumingFunction
	// go useFanOutChannel(sendChannel, fanOutChannel)

	go useThrottlingInFanOutChannel(sendChannel, fanOutChannel)

	// Reading Values from Fan Out Channel
	// This will return square of the values present in sendChannel
	for value := range fanOutChannel {
		fmt.Println("Values from Fan Out Channel: ", value)
	}

	fmt.Println("End of FanOut code block")

}

func populate(channel chan<- int) {
	for i := 0; i < 5; i++ {
		channel <- i
	}
	close(channel)
}

func useFanOutChannel(sendChannel <-chan int, fanOutChannel chan<- int) {
	var wg sync.WaitGroup
	for value := range sendChannel {
		wg.Add(1)
		go func(localValue int) {
			// Call some time consuming step(function) here
			fanOutChannel <- timeConsumingFunction(localValue)
			wg.Done()
		}(value)
	}
	wg.Wait()
	close(fanOutChannel)
}

func useThrottlingInFanOutChannel(sendChannel <-chan int, fanOutChannel chan<- int) {
	var wg sync.WaitGroup
	const goroutines = 2
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for value := range sendChannel {
				// Call some time consuming step(function) here
				fanOutChannel <- timeConsumingFunction(value)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(fanOutChannel)
}

func timeConsumingFunction(value int) int {
	// Returing back square of passed value and adding sleep
	// to pretend that this function takes time
	square := value * value
	time.Sleep(time.Microsecond)
	return square
}
