package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{name: "John", age: 30}
	fmt.Println(p1.age)
	p2 := person{"Alice", 25}
	fmt.Println(p2)

	p3 := person{name: "Bob"}
	fmt.Println(p3)

	p4 := person{age: 40}
	fmt.Println(p4)

	p5 := person{}
	fmt.Println(p5)

	p6, err := newPerson("Charlie", 35)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(p6.name, p6.age)
	}

	//create scan for user input and refer to the struct fields
	var person = person{}
	fmt.Println("Enter name:")
	fmt.Scanln(&person.name)
	fmt.Println("Enter age:")
	fmt.Scanln(&person.age)
	person2, err := newPerson(person.name, person.age)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Person2: %s, %d\n", person2.name, person2.age)
	}

}

// newPerson is a constructor function that creates and returns a pointer to a new person struct.
func newPerson(name string, age int) (*person, error) {
	//validation can be added here if needed
	if name == "" || age < 0 {
		return nil, fmt.Errorf("invalid person data")
	}

	return &person{
		name: name,
		age:  age,
	}, nil
}
