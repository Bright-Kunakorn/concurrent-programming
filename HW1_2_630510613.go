package main

import (
	"fmt"
	"time"
)

// PlayPing simulates the "bob" player in a game of ping-pong.
// It receives messages on the "ping" channel and sends responses on the "pong" channel.
func PlayPing(ping <-chan string, pong chan<- string) {
	// Continuously receive messages on the "ping" channel
	for msg := range ping {
		// Print the received message
		fmt.Println("bob: " + msg)
		// Sleep for one second
		time.Sleep(time.Second)
		// Send a "pong" message on the "pong" channel
		pong <- "pong"
	}
}

// PlayPong simulates the "alice" player in a game of ping-pong.
// It receives messages on the "pong" channel and sends responses on the "ping" channel.
func PlayPong(ping chan<- string, pong <-chan string) {
	// Continuously receive messages on the "pong" channel
	for msg := range pong {
		// Print the received message
		fmt.Println("alice: " + msg)
		// Sleep for one second
		time.Sleep(time.Second)
		// Send a "ping" message on the "ping" channel
		ping <- "ping"
	}
}

func main() {
	// Create the "alice" and "bob" channels
	alice := make(chan string)
	bob := make(chan string)

	// Run the PlayPing and PlayPong functions in separate goroutines
	go PlayPing(alice, bob)
	go PlayPong(alice, bob)

	// Send a "ping" message on the "alice" channel to start the game
	alice <- "ping"

	// Enter an infinite loop to keep the program running
	for {
	}
}
