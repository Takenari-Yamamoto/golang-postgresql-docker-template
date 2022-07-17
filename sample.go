package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		max int
		msg string
		x bool
	)


	// IntVar 整数オプション
	// flag.IntVar(&変数名, "オプションの名前", デフォルト値, "オプションの説明文")
	flag.IntVar(&max, "n", 32, "処理数の最大値")
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	flag.BoolVar(&x, "x", false, "拡張オプション")

	flag.Parse()

	fmt.Println(max)
	fmt.Println(msg)
	fmt.Println(x)
}
