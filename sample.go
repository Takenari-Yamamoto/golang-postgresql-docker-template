// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
}
