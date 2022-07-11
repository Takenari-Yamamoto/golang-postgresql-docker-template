// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"

// fmtモジュールをインポート

func main() {
	var s string = "Hello golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	fmt.Println(`test
		test
	test
	`)

	fmt.Println(string(s[0]))
}
