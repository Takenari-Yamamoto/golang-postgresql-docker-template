package main

import (
	"fmt"
	"hello-module/todo_app/config"
)

func main () {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)
}
