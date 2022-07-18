package main

import (
	"fmt"
	"hello-module/todo_app/app/models"
)

func main () {

	// ユーザー作成
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@test.jp"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// ユーザー取得
	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// ユーザー更新
	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// ユーザー削除
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// Todo作成
	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

	// Todo取得
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

		// Todo作成
	user, _ := models.GetUser(2)
	user.CreateTodo("Second Todo")

	todos, _ := models.GetTodos()
	fmt.Println(todos)
}
