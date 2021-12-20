package service

import (
	"ImGoDeveloper"
	"ImGoDeveloper/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const SALT = "fjklnaisfbibxso"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user ImGoDeveloper.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(SALT)))
}
