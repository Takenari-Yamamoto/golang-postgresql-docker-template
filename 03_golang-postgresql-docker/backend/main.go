package main

import (
	"fmt"
	"golang-postgresql-docker/database"
)

func main () {
	fmt.Println("Hello world")
	db, err := database.NewDb()
	if err != nil {
		fmt.Println("エラーだよ",err)
		return
	}

	fmt.Println("connect to success", db)
}
