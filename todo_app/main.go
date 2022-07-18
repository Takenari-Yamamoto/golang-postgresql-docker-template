package main

import (
	"fmt"
	"hello-module/todo_app/app/models"
)

func main () {

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@test.jp"
	u.Password = "testtest"
	fmt.Println(u)

	u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	user, _ := models.GetUser(2)
	user.CreateTodo("First Todo")

}
