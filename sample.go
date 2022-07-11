// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"

// fmtモジュールをインポート

func main() {
	var x interface {}
	x = 1
	fmt.Println(x)
}
