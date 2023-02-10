package repository

import (
	"context"
	"fmt"
	"golang-postgresql-docker/database"
	"golang-postgresql-docker/database/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TodoRepo struct {}

func NewTodoRepository() *TodoRepo {
	return &TodoRepo{}
}

func (repo TodoRepo) Insert(title, content string) {
	db, err := database.NewDb()
	if err != nil {
		fmt.Println("エラーだよ",err)
		return
	}

	uuidObj, _ := uuid.NewUUID()
	todo := models.Todo{
		ID: uuidObj.String(),
		Title: title,
		Content: content,
	}
	todo.Insert(context.Background(), db, boil.Infer())

}
