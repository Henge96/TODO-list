package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "todo-app"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (name, username, password_hash) values ($1, $2, $3) returning id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("row.Scan: %w", err)
	}

	return id, nil
}

func (r *AuthPostgres) GetUserByUsername(username string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("select * from %s where username=$1", userTable)

	err := r.db.Get(&user, query, username)
	return user, err
}
