package service

import (
	todo "todo-app"
	"todo-app/pkg/repo"
)

type AuthorizationService interface {
	CreateUser(user todo.User) (id int, err error)
	GenerateToken(username, password string) (token string, err error)
}

type TODOList interface {
}

type TODOItem interface {
}

type Service struct {
	AuthorizationService AuthorizationService
	TODOList             TODOList
	TODOItem             TODOItem
}

func New(repo *repo.Repo) *Service {
	return &Service{
		AuthorizationService: NewAuthService(repo.AuthorizationRepo),
	}
}
