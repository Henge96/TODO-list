package service

import (
	"crypto/sha1"
	"fmt"
	todo "todo-app"
	"todo-app/pkg/repo"
)

type AuthService struct {
	repo repo.AuthorizationRepo
}

func NewAuthService(repo repo.AuthorizationRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return "", fmt.Errorf("s.repo.GetUser: %w", err)
	}

}

func generatePasswordHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash)
}
