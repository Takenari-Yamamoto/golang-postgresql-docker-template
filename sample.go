package main

import "fmt"

func main() {
	// var ch1 chan int

	// ch1 = make(chan int)

	// fmt.Println(cap(ch1))
	// ch2 := make(chan int)
	// fmt.Println(cap(ch2))


	// データの操作
	// データの送受信を行うのがチャネル
	ch3 := make(chan int, 5)

	// ① データをチャネルに送信する
	// 下記の例は ch3 に 1というデータを送信している
	ch3 <- 1
	// データの要素数は len() で調べることができる
	fmt.Println(len(ch3)) // 1

	ch3 <- 2
	ch3 <- 3
	fmt.Println("長さ", len(ch3)) // 3

	i := <-ch3
	fmt.Println("長さ", len(ch3))
	fmt.Println(i)
	i2 := <-ch3
	fmt.Println("長さ", len(ch3))
	fmt.Println(i2)
	i3 := <-ch3
	fmt.Println("長さ", len(ch3))
	fmt.Println(i3)

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	// ch3 <- 6

}
