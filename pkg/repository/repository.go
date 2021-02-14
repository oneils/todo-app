package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/oneils/todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string, passwordHash string) (todo.User, error)
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

// New creates a new Repository
func New(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
