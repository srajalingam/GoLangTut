package main

import (
	"fmt"
	"moduleStructs/user"
)

func main() {
	u := user.User{}
	getUserInputFromScan(&u)
}

func getUserInputFromScan(u *user.User) {
	fmt.Println("Enter your name and age:")
	fmt.Scanln(&u.Name)
	fmt.Println("Enter your age:")
	fmt.Scanln(&u.Age)
	u.OutputUserDetails()
}
