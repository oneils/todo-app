package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/oneils/todo-app"
)

type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres creates a new AuthPostgres
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser saves a new user in Post
func (r AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// GetUser finds a user in DB according to the username and hashed password specified.
func (r AuthPostgres) GetUser(username string, passwordHash string) (todo.User, error) {
	var user todo.User

	query := fmt.Sprintf("SELECT id from %s WHERE username=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, passwordHash)

	return user, err
}
