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

func (repo *TodoRepo) FindById(id string) (*gql.Todo, error) {
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

	return convModelToGql(todo), nil
}


func (repo *TodoRepo) GetAll() ([]*gql.Todo, error) {
	db, err := database.NewDb()

	if err != nil {
		fmt.Println("エラーだよ",err)
		return nil, err
	}
	todos, err := models.Todos().All(context.Background(), db)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var gqlTodos []*gql.Todo
	for _, t := range todos {
		gqlTodos = append(gqlTodos, convModelToGql(t))
	}

	return gqlTodos, nil
}

func (repo *TodoRepo) Insert(ctx context.Context,title, content string) (*gql.Todo, error) {
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
	if err := todo.Insert(ctx, db, boil.Infer()); err != nil {
		fmt.Println("エラーだよ",err)
		return nil, err
	}

	return convModelToGql(&todo), nil
}

func convModelToGql(t *models.Todo) *gql.Todo {
	return &gql.Todo{
		ID: t.ID,
		Title: t.Title,
		Content: t.Content,
	}
}
