package service

import (
	"errors"

	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/configs"
	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/model"
	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func (s *AuthService) Register(username, password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := model.User{
		Username: username,
		Password: string(hash),
	}

	return s.Repo.Create(user)
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.Repo.GetByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := configs.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
