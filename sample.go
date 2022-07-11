// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"

// fmtモジュールをインポート

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n",arr1)

	var arr2 [3] string = [3]string{"A","B"}
	fmt.Println(arr2)

	arr3 := [3]int{1,2,3}
	fmt.Println(arr3)

	arr4 := [...]string{ "C", "D" }
	fmt.Println(arr4)

	fmt.Println(arr3[0])
}
