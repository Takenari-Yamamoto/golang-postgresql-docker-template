package main

import (
	"fmt"
	"hello-module/todo_app/app/controllers"
	"hello-module/todo_app/app/models"
)

func main () {

	fmt.Println(models.Db)

	controllers.StartMainServer()
}
