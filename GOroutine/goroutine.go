package main

import (
	"fmt"
	"time"
)

func greet(val string, doneChan chan bool) {
	fmt.Println("hi", val)
	doneChan <- true
}

func slowGreet(val string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("hello", val)
	doneChan <- true // <- pushing the value to channel
}

func main() {
	//done := make(chan bool)
	dones := make([]chan bool, 4)
	dones[0] = make(chan bool)
	go greet("Nice to meet you", dones[0])
	dones[1] = make(chan bool)
	go greet("how are you", dones[1])
	dones[2] = make(chan bool)
	go slowGreet("How .are ....you", dones[2])
	dones[3] = make(chan bool)
	go greet("hope you reached..", dones[3])
	// <-done // is done
	// <-done
	// <-done
	// <-done
	for _, done := range dones {
		<-done
	}
}
