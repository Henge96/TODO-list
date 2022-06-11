package repo

import (
	"github.com/jmoiron/sqlx"
	todo "todo-app"
)

type AuthorizationRepo interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string) (todo.User, error)
}

type TODOList interface {
}

type TODOItem interface {
}

type Repo struct {
	AuthorizationRepo AuthorizationRepo
	TODOList          TODOList
	TODOItem          TODOItem
}

func New(db *sqlx.DB) *Repo {
	return &Repo{
		AuthorizationRepo: NewAuthPostgres(db),
	}
}
