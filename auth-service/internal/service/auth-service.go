package service

import (
	"errors"
	"log"

	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/configs"
	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/model"
	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func (s *AuthService) Register(username, password string) error {
	log.Printf("Register attempt: username=%s", username)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("ERROR hashing password for user=%s: %v", username, err)
		return err
	}

	user := model.User{
		Username: username,
		Password: string(hash),
	}

	err = s.Repo.Create(user)
	if err != nil {
		log.Printf("ERROR creating user=%s: %v", username, err)
		return err
	}

	log.Printf("User registered successfully: username=%s", username)
	return nil
}

func (s *AuthService) Login(username, password string) (string, error) {
	log.Printf("Login attempt: username=%s", username)

	user, err := s.Repo.GetByUsername(username)
	if err != nil {
		log.Printf("ERROR user not found: username=%s", username)
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Printf("ERROR invalid password: username=%s", username)
		return "", errors.New("invalid credentials")
	}

	token, err := configs.GenerateToken(user.ID)
	if err != nil {
		log.Printf("ERROR generating token: userID=%d err=%v", user.ID, err)
		return "", err
	}

	log.Printf("Login successful: username=%s userID=%d", username, user.ID)
	return token, nil
}
