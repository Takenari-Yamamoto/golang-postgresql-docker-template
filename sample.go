// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"

func main() {
	// var sl []int
	// fmt.Println(sl)

	// var sl2 []int = []int{100, 200}
	// fmt.Println(sl2)
	// sl2[0] = 1000
	// fmt.Println(sl2)

	// sl3 := []string{"A", "B"}
	// fmt.Println(sl3)

	// sl4 := make([]int, 5)
	// fmt.Println(sl4)

	sl5 := []int {1,2,3,4,5}
	fmt.Println(sl5[0])
	fmt.Println(sl5[2:5])

}
