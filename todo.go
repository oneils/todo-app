package todo

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
	ListId int `json:"listId"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int `json:"id"`
	ListId int `json:"listId"`
	ItemId int `json:"itemId"`
}

type UpdateTodoListRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// ValidateNoUpdate validates if UpdateTodoListRequest updates or not. Return error when there were no changes for UpdateTodoListRequest.
func (i UpdateTodoListRequest) ValidateNoUpdate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
