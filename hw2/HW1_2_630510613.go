package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func play(alice string, bob string) {
	fmt.Println(alice + " : ping")
	time.Sleep(time.Millisecond * 100)
	fmt.Println(bob + " : pog")
	time.Sleep(time.Millisecond * 100)
	defer wg.Done()

}
func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go play("alice", "bob")
		wg.Wait()
	}

}
