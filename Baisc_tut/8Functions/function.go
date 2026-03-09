package main

import "fmt"

func main() {

	fmt.Println(add(2, 3))

	fmt.Println(operate(4, 5, add))

	fmt.Println(sub(5, 2))

	closureFunc := closure()
	closureFunc()

	closureWithParamFunc := closureWithParam()
	closureWithParamFunc(2)
	closureWithParamFunc(3)

	closureWithReturnFunc := closureWithReturn()
	fmt.Println(closureWithReturnFunc(2))
	fmt.Println(closureWithReturnFunc(3))

	fmt.Println(factorial(5))

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

//func to create recurion factorial
// bigo: O(n)
// space complexity: O(n) due to recursive call stack
//time complexity: O(n) due to recursive calls

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
