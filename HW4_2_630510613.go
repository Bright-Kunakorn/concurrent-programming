package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wakeBell chan bool
var sleepBell chan bool

func barberShop(chairs chan string, ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	sleepBell <- true
	defer wg.Done()
	fmt.Println("BarberShop opened")
	defer fmt.Println("BarberShop closed")
	closeShop := false
	for {
		select {
		case id := <-chairs:
			getHairCut(id)
			if len(chairs) == 0 {
				if !closeShop {
					sleepBarber(id)
				} else {
					return
				}
			}
		case <-wakeBell:
			wakeBarber()
		case <-ctx.Done():
			if len(chairs) == 0 {
				return
			}
			closeShop = true
		}
	}
}

func customerEnter(id int, chairs chan string) {
	fmt.Printf("Customer %d enter BarberShop\n", id)
	select {
	case <-sleepBell:
		wakeBell <- true
		chairs <- strconv.Itoa(id)
		fmt.Printf("customer %d sitting\n", id)
	case chairs <- strconv.Itoa(id):
		fmt.Printf("customer %d sitting\n", id)
	default:
		fmt.Printf("No chair available, customer %d leaving BarberShop\n", id)
	}
}

func sleepBarber(id string) {
	sleepBell <- true
	fmt.Println("Barber sleeping")
}
func wakeBarber() {
	fmt.Println("Wakeup barber")
}
func getHairCut(id string) {
	fmt.Println("Cutting for customer ", id)
	time.Sleep(1 * time.Second)
	fmt.Println("Cutting done, customer", id)
}

func main() {
	wakeBell = make(chan bool)
	sleepBell = make(chan bool, 1)
	chairs := make(chan string, 5)
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())
	go barberShop(chairs, ctx, wg)
	<-sleepBell
	for i := 1; i <= 10; i++ {
		customerEnter(i, chairs)
	}
	cancel()
	wg.Wait()
	close(wakeBell)
	close(chairs)
	fmt.Println("End")
}
