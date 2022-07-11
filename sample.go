// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"

// fmtモジュールをインポート

func main() {
	byteA := []byte{ 72, 73 }
	fmt.Println(byteA)

	// Byte to String
	fmt.Println(string(byteA))

	// String to Byte
	c := [] byte("HI")
	fmt.Println(c)
}
