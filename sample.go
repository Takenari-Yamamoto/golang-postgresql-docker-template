package main // mainパッケージであることを宣言

import "fmt" // fmtモジュールをインポート

func main() {		// 最初に実行されるmain()関数を定義

	// コンパイル時に個数が決まっていて変更不可なものを配列という（array）
	// a1 := [3]string{}
	// a1[0] = "ggg"
	// a1[1] = "ggggg"
	// a1[2] = "rwerwerwe"
	// fmt.Println(a1)

	// メモリ効率や速度は若干落ちますが、個数を変更可能なものを スライス と呼びます。
	// s1 := []string{}
	// s1 = append(s1, "Red")
	// s1 = append(s1, "Blue")
	// s1 = append(s1, "Yellow")
	// fmt.Println(len(s1))
	// fmt.Println(cap(s1))

	m1 := map[string]int {
		"x": 100,
		"y": 200,
	}

	fmt.Println(m1["x"])
}
