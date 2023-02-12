package graph

import (
	"golang-postgresql-docker/repository"
)

type Resolver struct{
	TodoRepository *repository.TodoRepo
}
