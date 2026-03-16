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
	done := make(chan bool)

	go greet("Nice to meet you", done)
	go greet("how are you", done)
	go slowGreet("How .are ....you", done)

	go greet("hope you reached..", done)
	<-done // is done
	<-done
	<-done
	<-done
}
