// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"


func main() {
	// var m = map[string]int {"A":100}
	// fmt.Println(m)

	// m2 := map[string]int {"A":100, "B": 200}
	// fmt.Println(m2)

	// m3 := map[int]string{
	// 	1: "A",
	// 	2: "B",
	// }
	// fmt.Println(m3)

	// m4 := make(map[int]string)
	// fmt.Println(m4)
	// m4[1] = "A"
	// m4[2] = "B"
	// // fmt.Println(m4)

	var m5 = map[string]string {"A":"HOGE"}
	fmt.Println(m5["A"])
	fmt.Println(m5["C"])

	s, ok := m5["A"]
	fmt.Println(s, ok)
}
