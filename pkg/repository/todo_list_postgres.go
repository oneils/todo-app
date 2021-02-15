package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/oneils/todo-app"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	var todoListId int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&todoListId); err != nil {
		tx.Rollback()
		return 0, err
	}

	usersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListTable)
	_, err = tx.Exec(usersListQuery, userId, todoListId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return todoListId, tx.Commit()
}

// FindAll finds all TodoList for specified userId
func (r TodoListPostgres) FindAll(userId int) ([]todo.TodoList, error) {
	var list []todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListTable)
	err := r.db.Select(&list, query, userId)
	return list, err
}

// FindById finds TodoList by specified listId
func (r TodoListPostgres) FindById(userId, id int) (todo.TodoList, error) {
	var todolist todo.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListTable)
	err := r.db.Get(&todolist, query, userId, id)

	return todolist, err
}

// Delete deletes Todolist for specified user
func (r TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListsTable, usersListTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

// Update updates the specified TodoList
func (r TodoListPostgres) Update(userId, listId int, updateListRequest todo.UpdateTodoListRequest) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateListRequest.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *updateListRequest.Title)
		argId++
	}

	if updateListRequest.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *updateListRequest.Description)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Infof("updateQuery: %s", query)
	logrus.Infof("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
