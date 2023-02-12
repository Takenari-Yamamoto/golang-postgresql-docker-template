package repository

import (
	"context"
	"fmt"
	"golang-postgresql-docker/database"
	"golang-postgresql-docker/database/models"
	gql "golang-postgresql-docker/graph/model"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TodoRepo struct {}

func NewTodoRepository() *TodoRepo {
	return &TodoRepo{}
}

func (repo *TodoRepo) FindById(id string) (*models.Todo, error) {
	db, err := database.NewDb()
	if err != nil {
		fmt.Println("エラーだよ",err)
		return nil, err
	}
	todo, err := models.Todos(
		models.TodoWhere.ID.EQ(id),
	).One(context.Background(), db)
	if err != nil {
		return nil, err
	}

	return todo, nil
}


func (repo *TodoRepo) GetAll() (models.TodoSlice, error) {
db, err := database.NewDb()
	if err != nil {
		fmt.Println("エラーだよ",err)
		return nil, err
	}
	todos, err := models.Todos().All(context.Background(), db)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (repo *TodoRepo) Insert(title, content string) (*models.Todo, error) {
	db, err := database.NewDb()
	if err != nil {
		fmt.Println("エラーだよ",err)
		return nil, err
	}

	uuidObj, _ := uuid.NewUUID()
	todo := models.Todo{
		ID: uuidObj.String(),
		Title: title,
		Content: content,
	}
	if err := todo.Insert(context.Background(), db, boil.Infer()); err != nil {
		return nil, err
	}

	return todo, nil
}

func convModelToGql(*models.Todo) *gql.Todo {

}
