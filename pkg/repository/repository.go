package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/oneils/todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
