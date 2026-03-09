package main

import "fmt"

func main() {

	fmt.Println(add(2, 3))

	fmt.Println(operate(4, 5, add))

	fmt.Println(sub(5, 2))

	closureFunc := closure()
	closureFunc()
}

// create function as value
func add(x, y int) int {
	return x + y
}

//function as type
type operation func(int, int) int

//use function as type
func operate(x, y int, op operation) int {
	return op(x, y)
}

//annonymous function
var sub = func(x, y int) int {
	return x - y
}

//closure function
func closure() func() {
	count := 0
	return func() {
		count++
		fmt.Println(count)
	}
}

//closure function with parameter
func closureWithParam() func(int) {
	count := 0
	return func(x int) {
		count += x
		fmt.Println(count)
	}
}

//closure function with return value
func closureWithReturn() func(int) int {
	count := 0
	return func(x int) int {
		count += x
		return count
	}
}
