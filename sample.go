// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"

func main() {
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 2)
	fmt.Println(i)

	i2 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i2)
}
