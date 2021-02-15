package service

import (
	"github.com/oneils/todo-app"
	"github.com/oneils/todo-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	FindById(userId, id int) (todo.TodoList, error)
	Delete(userId, id int) error
	Update(userId, id int, updateListRequest todo.UpdateTodoListRequest) error
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
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
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
