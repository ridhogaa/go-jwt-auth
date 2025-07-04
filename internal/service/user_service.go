package service

import (
	"fmt"

	"github.com/ridhogaa/go-jwt-auth/internal/model"
	"github.com/ridhogaa/go-jwt-auth/internal/repository"
	"github.com/ridhogaa/go-jwt-auth/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(username, password string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := model.User{
		Username: username,
		Password: string(hashedPassword),
	}
	return s.repo.CreateUser(user)
}

func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := util.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
