package main

import (
	"fmt"
	"hello-module/todo_app/app/models"
)

func main () {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@test.jp"
	u.Password = "testtest"
	fmt.Println(9999999, u)

	u.CreateUser()

}
