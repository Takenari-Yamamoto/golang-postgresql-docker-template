package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 構造体の例
	p := Person{"John Doe", 30}
	fmt.Println(p)

	// スライスの例
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// マップの例
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors)
}
