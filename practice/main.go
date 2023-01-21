package main

import "fmt"

type User struct {
	Name string
	Age int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
 	user.Age = 1000
}

func main() {
	var user1 User
	user1.Name = "Takenari"
	user1.Age = 25
	fmt.Println(user1)

	user2 := User{}
	user2.Name = "Miki"
	fmt.Println(user2)

	User3 := User{Name: "Shiba", Age: 25}
	fmt.Println(User3)

	user4 := User{"Yuma", 25}
	fmt.Println(user4)

	// new を使うと構造体のポインタ型を返す
	user5 := new(User)
	fmt.Println(user5)

	// アドレス演算子を使っても構造体のポインタ型を返す
	user6 := &User{}
	fmt.Println(user6)

	UpdateUser(user1)
	UpdateUser2(user6)

	// fmt.Println(user1)
	fmt.Println(user1)
}
