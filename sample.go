package main

import "fmt"

type User struct {
	Name string
	Age int
}

// レシーバー宣言
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SayAge() {
	fmt.Println(u.Age)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	user1.SetName("A")
	user1.SayName()

	user1.SetName2("HOGEHOGEHOGE")
	user1.SayName()

	user2 := &User{Name: "user2"}
	user2.SetName("BB")
	user2.SayName()
}
