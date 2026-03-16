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
	close(doneChan)
}

func main() {
	done := make(chan bool)
	//dones := make([]chan bool, 4)
	// dones[0] = make(chan bool)
	go greet("Nice to meet you", done)
	// dones[1] = make(chan bool)
	go greet("how are you", done)
	// dones[2] = make(chan bool)
	go slowGreet("How .are ....you", done)
	// dones[3] = make(chan bool)
	go greet("hope you reached..", done)
	// <-done // is done
	// <-done
	// <-done
	// <-done
	// for _, done := range dones {
	// 	<-done
	// }

	for range done {

	}
}
