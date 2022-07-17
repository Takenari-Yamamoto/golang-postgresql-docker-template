package main

import (
	"fmt"
	"hello-module/todo_app/config"
	"log"
)

func main () {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
