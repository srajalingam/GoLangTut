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
}
