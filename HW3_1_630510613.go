package main
import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var buffer = make([]byte, 0, 10)
var all sync.WaitGroup

var bufferChan = make(chan bool, 1)

func main() {
	rand.Seed(10)
	all.Add(1)
	go writer('a')
	go writer('b')
	go consumer()
	all.Wait()
	
}

func writer(c byte) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Int63n(1e9)))
		bufferChan <- true
		lb := len(buffer)
		if lb < cap(buffer) {
			buffer = buffer[:lb+1]
			buffer[lb] = c
			fmt.Printf("'%c' written to buffer.     buffer contents: %s\n",
				c, string(buffer))
		}
		<-bufferChan
	}
}

func consumer() {
	a := []byte{'a'}
	b := []byte{'b'}
	for i := 0; i < 5; {
		time.Sleep(time.Duration(rand.Int63n(1e9)))
		bufferChan <- true
		ai := bytes.Index(buffer, a)
		bi := bytes.Index(buffer, b)
		if ai >= 0 && bi >= 0 {
			if ai > bi {
				ai, bi = bi, ai
			}
			copy(buffer[bi:], buffer[bi+1:])
			copy(buffer[ai:], buffer[ai+1:])
			buffer = buffer[:len(buffer)-2]
			fmt.Printf("pair removed from buffer.  buffer contents: %s\n",
				string(buffer))
			i++
		}
		<-bufferChan
	}
	all.Done()
}
