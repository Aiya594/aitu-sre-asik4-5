package service

import (
	"errors"
	"log"
	"strings"

	"github.com/Aiya594/aitu-sre-asik4-5-user-profile/internal/model"
	repository "github.com/Aiya594/aitu-sre-asik4-5-user-profile/internal/repositoy"
)

type ProfileService struct {
	Repo *repository.ProfileRepository
}

func (s *ProfileService) CreateProfile(
	userID int,
	email string,
	phone string,
	address string,
) error {

	log.Printf(
		"[ProfileService] CreateProfile user=%d",
		userID,
	)

	if userID <= 0 {
		return errors.New("invalid user id")
	}

	if !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}

	profile := model.UserProfile{
		UserID:  userID,
		Email:   email,
		Phone:   phone,
		Address: address,
	}

	return s.Repo.Create(profile)
}

func (s *ProfileService) GetProfile(
	userID int,
) (*model.UserProfile, error) {

	log.Printf(
		"[ProfileService] GetProfile user=%d",
		userID,
	)

	return s.Repo.GetByUserID(userID)
}

func (s *ProfileService) UpdateProfile(
	userID int,
	email string,
	phone string,
	address string,
) error {

	log.Printf(
		"[ProfileService] UpdateProfile user=%d",
		userID,
	)

	if !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}

	profile := model.UserProfile{
		UserID:  userID,
		Email:   email,
		Phone:   phone,
		Address: address,
	}

	return s.Repo.Update(profile)
}
