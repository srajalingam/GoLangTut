package main

import (
	"fmt"
	"time"
)

func greet(val string) {
	fmt.Println("hi", val)
}

func slowGreet(val string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("hello", val)
	doneChan <- true // <- pushing the value to channel
}

func main() {
	go greet("Nice to meet you")
	go greet("how are you")
	done := make(chan bool)
	go slowGreet("How .are ....you", done)

	go greet("hope you reached..")
	<-done // is done
}
