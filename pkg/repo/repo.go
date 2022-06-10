package repo

import "github.com/jmoiron/sqlx"

type Authorisation interface {
}

type TODOList interface {
}

type TODOItem interface {
}

type Repo struct {
	Authorisation
	TODOList
	TODOItem
}

func New(db *sqlx.DB) *Repo {
	return &Repo{}
}
