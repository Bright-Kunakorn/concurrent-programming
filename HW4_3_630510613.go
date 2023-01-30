package main

import (
	"fmt"
	"sync"
	"time"
)
const (
	numBees = 10
	potSize = 10
)

type HoneyPot struct {
	honey int
	mux   sync.Mutex
}

func (pot *HoneyPot) addHoney() {
	pot.mux.Lock()
	defer pot.mux.Unlock()
	if pot.honey < potSize {
		pot.honey++
		if pot.honey == potSize {
			fmt.Println("Waking up bear")
		}
	}
}

func (pot *HoneyPot) eatHoney() {
	pot.mux.Lock()
	defer pot.mux.Unlock()
	pot.honey = 0
}

func (pot *HoneyPot) isEmpty() bool {
	pot.mux.Lock()
	defer pot.mux.Unlock()
	return pot.honey == 0
}

func (pot *HoneyPot) isFull() bool {
	pot.mux.Lock()
	defer pot.mux.Unlock()
	return pot.honey >= potSize
}

func main() {
	pot := HoneyPot{}
	done := make(chan bool)

	for {
		for i := 0; i < numBees; i++ {
			go func() {
				for {
					pot.addHoney()
					if pot.isFull() {
						done <- true
						return
					}
					fmt.Println("Bee is gathering honey")
					time.Sleep(500 * time.Millisecond)
				}
			}()
		}
		<-done
		for !pot.isEmpty() {
			pot.eatHoney()
			fmt.Println("Bear is eating honey")
			fmt.Println("Bear is resting")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
