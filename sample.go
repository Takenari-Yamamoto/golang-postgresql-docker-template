// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"

// fmtモジュールをインポート

// 頭文字を大文字にすることで他のパッケージからも呼び出せる
const Pi = 62

const (
	URL = "http://xxx.com"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

const Big = 913413413413413453 + 1

func main() {
	fmt.Println(Pi)

	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
}
