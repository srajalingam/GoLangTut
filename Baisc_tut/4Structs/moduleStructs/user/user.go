package user

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.Age, u.Name)
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}
