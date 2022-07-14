// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import (
	"fmt"
	"os"
)

// func TestDefer() {
// 	defer fmt.Println("END")
// 	fmt.Println("START")
// }

func RunDefer () {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func main() {
	RunDefer()

	file, err := os.Create("text.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("Hello"))
}
