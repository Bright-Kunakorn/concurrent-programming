package main

import (
	"fmt"
	"sync"
)

const maxCustomers = 3

var (
	customers     int
	customerMutex sync.Mutex
	customer      = make(chan struct{})
	barber        = make(chan struct{})
	customerDone  = make(chan struct{})
	barberDone    = make(chan struct{})
)

func main() {
	fmt.Print("")
	go barberThread()
	for i := 1; i <= 5; i++ {
		go customerThread(i)
	}
}

func customerThread(id int) {
	customerMutex.Lock()
	if customers == maxCustomers {
		customerMutex.Unlock()
		balk()
	}
	customers++
	customerMutex.Unlock()

	fmt.Println("Customer", id, "entering shop")
	customer <- struct{}{}
	<-barber
	getHairCut()
	customerDone <- struct{}{}
	<-barberDone
	fmt.Println("Customer", id, "leaving shop")
}

func barberThread() {
	for {
		<-customer
		cutHair()
		barber <- struct{}{}
		<-customerDone
		barberDone <- struct{}{}
	}
}

func getHairCut() {
	fmt.Println("Getting hair cut")
}

func cutHair() {
	fmt.Println("Cutting hair")
}

func balk() {
	fmt.Println("Shop is full, customer unable to enter")
}
