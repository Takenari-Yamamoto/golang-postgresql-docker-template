// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"

// func ReturnFunc() func() {
// 	return func() {
// 		fmt.Println("I am a function")
// 	}
// }

func CallFunction (f func()) {
	f()
}

func main() {
	// f := ReturnFunc()
	// f()
	CallFunction(func() {
		fmt.Println("I AM A FUNC")
	})
}
