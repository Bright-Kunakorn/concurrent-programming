package main

import (
	"fmt"
	"time"
)

const (
	numConsumers = 2
	numProducers = 4
	numMessages  = 6
)

type message struct {
	str  string
	wait chan bool
}

func producer(id int, ch chan<- *message) {
	for i := 0; i < numMessages; i++ {
		m := &message{
			fmt.Sprintf("message %d from producer %d", i, id),
			make(chan bool),
		}
		ch <- m
		fmt.Println("Sent message", m.str)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(id int, ch <-chan *message) {
	for m := range ch {
		fmt.Println("Received message", m.str, "in consumer", id)
		time.Sleep(50 * time.Millisecond)
		m.wait <- true
	}
}

func main() {
	ch := make(chan *message)

	for i := 0; i < numConsumers; i++ {
		go consumer(i, ch)
	}

	for i := 0; i < numProducers; i++ {
		go producer(i, ch)
	}

	// Wait for all the messages to be processed
	for i := 0; i < numMessages*numProducers; i++ {
		<-ch.Wait
	}
}
