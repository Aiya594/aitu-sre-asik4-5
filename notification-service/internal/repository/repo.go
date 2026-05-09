package repository

import (
	"database/sql"

	"github.com/Aiya594/aitu-sre-asik4-5-notification/internal/model"
)

type NotificationRepository struct {
	DB *sql.DB
}

func (r *NotificationRepository) Create(
	notification model.Notification,
) error {

	_, err := r.DB.Exec(
		`
		INSERT INTO notifications (
			user_id,
			message,
			type
		)
		VALUES ($1, $2, $3)
		`,
		notification.UserID,
		notification.Message,
		notification.Type,
	)

	return err
}
