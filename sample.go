package main

import "fmt"

type User struct {
	Name string
	Age int
}

func UpdateUser (user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2 (user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User

	user1.Name = "Takenari"
	user1.Age = 24
	fmt.Println(user1)

	user2 := User{}
	user2.Name = "Bob"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 34}
	fmt.Println(user3)

	user4 := User{"Sake", 20}
	fmt.Println(user4)

	user5 := User{ "user5", 300}
	fmt.Println(user5)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	// ポインタ型で変数定義する方法①
	user7 := new(User)
	fmt.Println(user7)

	// ポインタ型で変数定義する方法②
	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	UpdateUser2(user8)
	fmt.Println(user1)
	fmt.Println(user8)

}
