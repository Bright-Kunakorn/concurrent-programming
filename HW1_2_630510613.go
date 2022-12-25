package main

import (
	"fmt"
	"time"
)

func PlayPing(ping <-chan string, pong chan<- string) {
	for msg := range ping {
		fmt.Println("bob:" + msg)
		time.Sleep(time.Second)
		pong <- "pong"
	}
}
func PlayPong(ping chan<- string, pong <-chan string) {
	for msg := range pong {
		fmt.Println("alice: " + msg)
		time.Sleep(time.Second)
		ping <- "ping"
	}
}
func main() {
	alice := make(chan string)
	bob := make(chan string)
	go PlayPing(alice, bob)
	go PlayPong(alice, bob)
	alice <- "ping"
	for {
	}
}
