// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import "fmt"


func main() {
	a := 1
	if a == 2 {
		fmt.Println("TWO")
	} else if a == 1 {
		fmt.Println("ONE")
	} else {
		fmt.Println("I dont know")
	}

	if b := 100; b == 109 {
		fmt.Println("one hundred")
	}

}
