package main

import "fmt"

type User struct {
	Name string
	Age int
}

func main() {
	m := map[int]User{
		1: {Name:"Takenari", Age: 25},
		2: {Name:"Taro", Age: 5},
		3: {Name:"Ti", Age: 25},
		4: {Name:"nari", Age: 25},
	}

	fmt.Println(m)

	m2 := map[User]string{
		{Name:"Takenari", Age: 25}: "Tokyo",
		{Name:"Ti", Age: 25}: "LA",
	}

	fmt.Println(m2)

	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "user3"}
	fmt.Println(m3)

	for _, v := range m3 {
		fmt.Println(v)
	}

}
