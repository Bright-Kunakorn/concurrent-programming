package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func say(s string) {
	fmt.Println(s)
	time.Sleep(time.Millisecond * 100)
	wg.Done()
}
func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("main : hello")
		wg.Add(1)
		go say("goroutine : world!")
		wg.Wait()
	}
}
