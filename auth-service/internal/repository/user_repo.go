package repository

import (
	"database/sql"

	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(user model.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		user.Username, user.Password,
	)
	return err
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	row := r.DB.QueryRow(
		"SELECT id, username, password FROM users WHERE username=$1",
		username,
	)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
