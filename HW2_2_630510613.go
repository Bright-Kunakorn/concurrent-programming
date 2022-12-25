package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool, 5)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go count(i, ch, &wg)
	}
	wg.Wait()
	fmt.Println("Final Sum:", counter)
}

func count(i int, ch chan bool, wg *sync.WaitGroup) {
	mu.Lock()
	for j := 1; j <= 10000; j++ {
		counter++
	}
	mu.Unlock()
	fmt.Println("From ", i, ":", counter)
	ch <- true
	wg.Done()
}
