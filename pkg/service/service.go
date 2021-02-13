package service

import "github.com/oneils/todo-app/pkg/repository"

type Authorization interface {
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

func New(repository *repository.Repository) *Service {
	return &Service{}
}
