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

	// Start 5 goroutines, each of which runs the count function.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go count(i)
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	// Print the final value of the counter variable.
	fmt.Println("Final Sum:", counter)
}

// count increments the counter variable 10000 times.
func count(i int) {
	// Decrement the WaitGroup counter when the function returns.
	defer wg.Done()

	// Increment the counter variable in a loop.
	for j := 1; j <= 10000; j++ {
		// Lock the mutex to synchronize access to the counter variable.
		mu.Lock()
		counter++
		// Unlock the mutex.
		mu.Unlock()
	}

	// Print the current value of the counter variable.
	fmt.Println("From ", i, ":", counter)
}
