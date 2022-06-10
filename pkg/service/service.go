package service

import "todo-app/pkg/repo"

type Authorisation interface {
}

type TODOList interface {
}

type TODOItem interface {
}

type Service struct {
	Authorisation
	TODOList
	TODOItem
}

func New(repo *repo.Repo) *Service {
	return &Service{}
}
