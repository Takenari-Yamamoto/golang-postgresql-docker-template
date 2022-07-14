// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"

func init() {
	fmt.Println("init")
}
func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}
