package main

import (
	"fmt"
	"sync"
)
// counter is a shared variable that will be incremented by multiple goroutines.
var counter int

// mu is a mutex that will be used to synchronize access to the counter variable.
var mu sync.Mutex

func main() {
	// Create a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// Create a channel with a capacity of 5.
	ch := make(chan bool, 5)

	// Start 5 goroutines, each of which runs the count function.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go count(i, ch, &wg)
	}

	// Wait for 5 values to be received on the channel.
	for i := 0; i < 5; i++ {
		<-ch
	}

	// Print the final value of the counter variable.
	fmt.Println("Final Sum:", counter)
}

// count increments the counter variable 10000 times.
func count(i int, ch chan bool, wg *sync.WaitGroup) {
	// Lock the mutex to synchronize access to the counter variable.
	mu.Lock()
	for j := 1; j <= 10000; j++ {
		counter++
	}
	// Unlock the mutex.
	mu.Unlock()

	// Print the current value of the counter variable.
	fmt.Println("From ", i, ":", counter)

	// Send a value on the channel to indicate that the function has finished.
	ch <- true

	// Decrement the WaitGroup counter.
	wg.Done()
}
