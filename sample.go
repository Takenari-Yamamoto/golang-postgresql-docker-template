// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import (
	"fmt"
	"strconv"
)

// fmtモジュールをインポート

func main() {
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3 = %T\n", i3)
	// fmt.Println(i3)

	var s string = "100"
	i, e := strconv.Atoi(s)
	fmt.Println(i, e)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("i = %T\n", e)
}
