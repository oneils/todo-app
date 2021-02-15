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

// GetAll return all Todolists for specified userId
func (s TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.FindAll(userId)
}

// FindById return all Todolists for specified userId
func (s TodoListService) FindById(userId, id int) (todo.TodoList, error) {
	return s.repo.FindById(userId, id)
}

// Delete deletes the list for specified user
func (s TodoListService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}

// Update updates the specified TodoList
func (s TodoListService) Update(userId, id int, updateListRequest todo.UpdateTodoListRequest) error {

	return s.repo.Update(userId, id, updateListRequest)
}
