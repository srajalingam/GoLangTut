package main

import (
	"fmt"
)

// create two strcut employee and manager and embed employee struct in manager struct
type employee struct {
	name   string
	age    int
	salary float64
}

type manager struct {
	employee   //embedding employee struct in manager struct
	department string
}

func main() {
	m1 := manager{
		employee:   employee{name: "John", age: 30, salary: 50000},
		department: "Sales",
	}
	fmt.Println(m1.name)
	fmt.Println(m1.age)
	fmt.Println(m1.salary)
	fmt.Println(m1.department)
}
