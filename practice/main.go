package main

import "fmt"

type User struct {
	Name string
	Age int
}

type Users []*User

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := User{Name: "User1", Age: 10}
	user2 := User{Name: "User2", Age: 20}
	user3 := User{Name: "User3", Age: 30}
	user4 := User{Name: "User4", Age: 40}

	users := Users{}

	users = append(users, &user1, &user2, &user3, &user4)

	fmt.Println(users)

	for _, u := range users {
		fmt.Println(u)
	}

	users2 := make([]*User, 0)
	users2 = append(users, &user1, &user2, &user3, &user4)

	for _, u := range users2 {
		fmt.Println(*u)
	}
}
