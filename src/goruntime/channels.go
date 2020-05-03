package goruntime

import (
	"fmt"
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
