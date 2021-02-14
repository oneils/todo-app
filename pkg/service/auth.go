package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/oneils/todo-app"
	"github.com/oneils/todo-app/pkg/repository"
)

const salt = "jlsdf9s0fd#@$_fsdfsdf#@SDF!23"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
