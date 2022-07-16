package main

import "fmt"

func Double(i int) {
	i = i * 2
}

func DoubleV2(i *int) {
	*i = *i * 2
}

func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)

	Double(n)

	fmt.Println(n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	*p = 6000
	fmt.Println(n)

	DoubleV2(&n)
	fmt.Println(n)

}
