package main

import "fmt"

type T struct {
	User User
}

type User struct {
	Name string
	Age int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "Take"}
	user1.SayName()

	user1.SetName("SAKE")
	user1.SayName()
}
