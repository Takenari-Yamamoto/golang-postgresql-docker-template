// mainパッケージであることを宣言
// ちなみにパッケージの宣言は 1つまでしかできない
package main

import (
	"fmt"
)


func main() {
	// i := 0
	// for  {
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
