package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/oneils/todo-app"
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
