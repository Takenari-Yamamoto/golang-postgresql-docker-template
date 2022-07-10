// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

// fmtモジュールをインポート
import "fmt"

func outer() {
	var s4 string = "outerやで！"
	fmt.Println(s4)
}

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "Hello World"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 100
		a2 string = "Golang"
	)

	fmt.Println(i2, a2)

	// 初期値がない時は0が返る
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "GOGOGOGO"
	fmt.Println(i3, s3)

	i4 := 666
	fmt.Println(i4)

	i4 = 342134
	fmt.Println(i4)

	outer()

}
