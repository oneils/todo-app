package service

import (
	"github.com/oneils/todo-app"
	"github.com/oneils/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

// NewTodoListService creates a new instance of TodoListService
func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

// Create creates a new TodoList for specified userId
func (s TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
