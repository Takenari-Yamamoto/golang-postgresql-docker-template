package main

import "fmt"

type T struct {
	User
}

type User struct {
	Name string
	Age int
}

func NewUser (name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("user1", 10)
	fmt.Println(user1)
	fmt.Println(*user1)
}
