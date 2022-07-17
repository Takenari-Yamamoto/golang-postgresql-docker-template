package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("かいし")
	time.Sleep(2 * time.Second)
	fmt.Println("おしまい")
	ch <- "実行結果"
}

func main() {
	// チャネルの作成
	ch := make(chan string)

	// context の作成
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()

	go longProcess(ctx, ch)

	L:
		for {
			select {
				case <- ctx.Done():
					fmt.Println("########Error########")
					fmt.Println(ctx.Err())
					break L
				case s := <-ch:
					fmt.Println(s)
					fmt.Println("success")
					break L
			}
		}

		fmt.Println("LOOP FIN")
}
