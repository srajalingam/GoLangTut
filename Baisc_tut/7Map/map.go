package main

import "fmt"

func main() {
	// map is a collection of key value pairs
	// map is unordered
	// map is mutable
	// map is indexed by keys

	websites := map[string]string{
		"google":   "www.google.com",
		"facebook": "www.facebook.com",
		"twitter":  "www.twitter.com",
	}
	fmt.Println(websites)

	// adding new key value pair to map
	websites["linkedin"] = "www.linkedin.com"
	fmt.Println(websites)

	// deleting a key value pair from map
	delete(websites, "twitter")
	fmt.Println(websites)

	// make function is used to create a map

	usernames := make(map[string]string)
	usernames["john"] = "john123"
	usernames["jane"] = "jane123"
	fmt.Println(usernames)

	scores := make([]string, 2, 4) // make function is used to allocate memory for slice
	scores[0] = "100"
	scores[1] = "90"
	scores = append(scores, "80") // append function is used to add new element to slice
	scores = append(scores, "70")
	scores = append(scores, "60") // when we append more than the capacity of the slice, it will create a new slice with double the capacity and copy the old elements to the new slice
	fmt.Println(scores)

	// iterating over map
	for key, value := range websites {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

}
