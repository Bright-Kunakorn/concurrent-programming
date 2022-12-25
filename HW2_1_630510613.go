package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go count(i)
	}
	wg.Wait()
	fmt.Println("Final Sum:", counter)
}
func count(i int) {
	defer wg.Done()
	for j := 1; j <= 10000; j++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
	fmt.Println("From ", i, ":", counter)
}
