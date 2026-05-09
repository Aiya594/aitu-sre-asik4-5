package service

import (
	"log"

	"github.com/Aiya594/aitu-sre-asik4-5-notification/internal/model"
	"github.com/Aiya594/aitu-sre-asik4-5-notification/internal/repository"
)

type NotificationService struct {
	Repo *repository.NotificationRepository
}

func (s *NotificationService) Send(
	userID int,
	message string,
	notifType string,
) error {

	log.Printf(
		"[NotificationService] user=%d type=%s",
		userID,
		notifType,
	)

	/*
		SIMULATION:
		Here later you can integrate:
		- SMTP
		- Telegram
		- Twilio
		- Firebase
	*/

	log.Printf(
		"[EMAIL SIMULATION] To user=%d: %s",
		userID,
		message,
	)

	notification := model.Notification{
		UserID:  userID,
		Message: message,
		Type:    notifType,
	}

	return s.Repo.Create(notification)
}
