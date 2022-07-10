// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"

// fmtモジュールをインポート

func main() {

		var i string = "100"

		var i2 int64 = 200

		fmt.Printf("%T\n", i)


		fmt.Printf("%T\n", int32(i2))
}
