package goruntime

import (
	"fmt"
)

// Channel demonstrates use of channels in go routines
func Channel() {
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
	readChannel := make(<-chan int)
	writeChannel := make(chan<- int)

	fmt.Printf("Read Channel: %T\n", readChannel)
	fmt.Printf("Write Channel: %T\n", writeChannel)
}

// DirectionalChannels demonstrates use of directional channels in Go
func DirectionalChannels() {
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
