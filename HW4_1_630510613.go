package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numSmokers = 3
	numAgents  = 1
)

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

func agent(tobacco, paper, matches chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(1000 * time.Millisecond)
		ingredient := rand.Intn(numSmokers)
		if ingredient == 0 {
			tobacco <- true
			paper <- true
			fmt.Println("Agent puts tobacco and paper on the table")
		} else if ingredient == 1 {
			paper <- true
			matches <- true
			fmt.Println("Agent puts paper and matches on the table")
		} else {
			tobacco <- true
			matches <- true
			fmt.Println("Agent puts tobacco and matches on the table")
		}
	}
}

func main() {
	var wg sync.WaitGroup
	tobacco := make(chan bool)
	paper := make(chan bool)
	matches := make(chan bool)

	wg.Add(numSmokers + numAgents)

	for i := 0; i < numSmokers; i++ {
		go smoker(i, tobacco, paper, matches, &wg)
	}

	for i := 0; i < numAgents; i++ {
		go agent(tobacco, paper, matches, &wg)
	}

	wg.Wait()
}
