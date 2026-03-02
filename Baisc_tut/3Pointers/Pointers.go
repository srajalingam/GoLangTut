package main

import "fmt"

func main() {
	age := 25
	fmt.Println("Age:", age)

	// Create a pointer to the age variable
	agePtr := &age
	fmt.Println("Pointer to Age:", agePtr)

	// Access the value using the pointer
	fmt.Println("Value at Pointer:", *agePtr)

	adultYears := getAdultYears(agePtr)
	fmt.Println("Adult Years:", adultYears)

}

func getAdultYears(age *int) int {
	if *age >= 18 {
		return *age - 18
	}
	return 0
}
