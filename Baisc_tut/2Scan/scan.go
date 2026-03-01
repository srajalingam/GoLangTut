package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %s!\n", name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Println("You are", age, "years old.")
}
