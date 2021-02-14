package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/oneils/todo-app"
	"github.com/oneils/todo-app/pkg/repository"
	"time"
)

const (
	salt       = "Hcfq3*CHGBGA=n,AsP,npfkfC]vCB*Ht.@EhFe~%"
	tokenTTL   = 12 * time.Hour
	signingKey = "TgWa!UEsrE+RYGENYG@VyRk0Qa-:pJGDiWZ_LQ2t"
)

type AuthService struct {
	repo repository.Authorization
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser generates a hashed password and creates a new user
func (a AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type tokenClaim struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// GenerateToken generates a new JWT token if a user found for specified username and password
func (a AuthService) GenerateToken(username string, password string) (string, error) {

	user, err := a.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}
