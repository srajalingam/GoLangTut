package main

import (
	"fmt"
	"time"
)

func greet(val string) {
	fmt.Println("hi", val)
}

func slowGreet(val string) {
	time.Sleep(3 * time.Second)
	fmt.Println("hello", val)
}

func main() {
	go greet("Nice to meet you")
	go greet("how are you")
	go slowGreet("How .are ....you")
	go greet("hope you reached..")
}
