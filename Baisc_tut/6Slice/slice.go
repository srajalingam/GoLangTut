package main

import "fmt"

func main() {
	// Create a slice of integers
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)
	// Append a new element to the slice
	numbers = append(numbers, 6)
	fmt.Println("After appending 6:", numbers)
	// Create a new slice that is a subset of the original slice
	subset := numbers[1:4]
	fmt.Println("Subset of the slice (index 1 to 3):", subset)
	// Modify the original slice and see how it affects the subset
	numbers[2] = 10
	fmt.Println("After modifying the original slice:", numbers)
	fmt.Println("Subset after modifying the original slice:", subset)

	//cap

	hobbies := []string{"Reading", "Traveling", "Cooking"}
	fmt.Println("Original hobbies slice:", hobbies)
	fmt.Printf("Length: %d, Capacity: %d\n", len(hobbies), cap(hobbies))

	mainHobbies := hobbies[:2]
	fmt.Println("Main hobbies slice:2", hobbies)
	fmt.Println("Main hobbies slice:", mainHobbies)
	fmt.Printf("Length: %d, Capacity: %d\n", len(mainHobbies), cap(mainHobbies))

	hobbies = hobbies[:2]
	fmt.Println("Hobbies slice after slicing to main hobbies:", hobbies)
	fmt.Printf("Length: %d, Capacity: %d\n", len(hobbies), cap(hobbies))

	disccountPrices := []float64{19.99, 29.99, 39.99, 49.99}
	prices := []float64{10, 20.0}

	prices = append(prices, disccountPrices...)
	fmt.Println("Prices after appending discount prices:", prices)

}
