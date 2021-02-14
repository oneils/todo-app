package service

import (
	"github.com/oneils/todo-app"
	"github.com/oneils/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService creates a new instance of Service
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
