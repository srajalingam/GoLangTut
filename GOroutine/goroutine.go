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
	greet("Nice to meet you")
	greet("how are you")
	slowGreet("How .are ....you")
	greet("hope you reached..")
}
