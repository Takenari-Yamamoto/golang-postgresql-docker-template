// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"
func main() {
	// sl = append(sl, 300)
	// // fmt.Println(sl)

	// sl = append(sl, 400, 500, 600)
	sl := []int{100, 200, 300, 400, 500}
	fmt.Println(cap(sl))
	fmt.Println(len(sl))

	sl2 := make([]int, 5, 10)
	fmt.Println(cap(sl2))
	fmt.Println(len(sl2))
}
