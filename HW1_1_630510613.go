package main

import (
	"fmt"
	"sync"
	"time"
)

// Declare a WaitGroup variable
var wg sync.WaitGroup

// Declare a function that takes a string argument
func say(s string) {
	// Print the string
	fmt.Println(s)
	// Sleep for 100 milliseconds
	time.Sleep(time.Millisecond * 100)
	// Decrement the WaitGroup counter by 1
	wg.Done()
}

func main() {
	// Run a loop 5 times
	for i := 0; i < 5; i++ {
		// Print a message
		fmt.Println("main : hello")
		// Increment the WaitGroup counter by 1
		wg.Add(1)
		// Start a new goroutine that calls the say function
		go say("goroutine : world!")
		// Block the main function until the WaitGroup counter is zero
		wg.Wait()
	}
}
