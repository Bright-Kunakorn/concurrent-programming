package main

import (
	"fmt"
	"sync"
	"time"
)
const (
	Smokers = 4
	Agents  = 1
)
func agent(tobacco, paper, matches chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(1000 * time.Millisecond)
		material := (Smokers)
		if material == 0 {
			tobacco <- true
			paper <- true
			fmt.Println("Agent puts tobacco and paper")
		} else if material == 1 {
			paper <- true
			matches <- true
			fmt.Println("Agent puts paper and matches")
		} else {
			tobacco <- true
			matches <- true
			fmt.Println("Agent puts tobacco and matches")
		}
	}
}
func smoker(id int, tobacco, paper, matches chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-tobacco:
			fmt.Printf("Smoker %d is smoking\n", id)
			time.Sleep(1000 * time.Millisecond)
		case <-paper:
			fmt.Printf("Smoker %d is smoking\n", id)
			time.Sleep(1000 * time.Millisecond)
		case <-matches:
			fmt.Printf("Smoker %d is smoking\n", id)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	tobacco := make(chan bool)
	paper := make(chan bool)
	matches := make(chan bool)
	wg.Add(Smokers + Agents)
	for i := 0; i < Smokers; i++ {
		go smoker(i, tobacco, paper, matches, &wg)
	}
	for i := 0; i < Agents; i++ {
		go agent(tobacco, paper, matches, &wg)
	}
	wg.Wait()
}
