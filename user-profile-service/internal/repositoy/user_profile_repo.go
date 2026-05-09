package repository

import (
	"database/sql"

	"github.com/Aiya594/aitu-sre-asik4-5-user-profile/internal/model"
)

type ProfileRepository struct {
	DB *sql.DB
}

func (r *ProfileRepository) Create(
	profile model.UserProfile,
) error {

	_, err := r.DB.Exec(
		`
		INSERT INTO user_profiles (
			user_id,
			email,
			phone,
			address
		)
		VALUES ($1, $2, $3, $4)
		`,
		profile.UserID,
		profile.Email,
		profile.Phone,
		profile.Address,
	)

	return err
}

func (r *ProfileRepository) GetByUserID(
	userID int,
) (*model.UserProfile, error) {

	row := r.DB.QueryRow(
		`
		SELECT
			id,
			user_id,
			email,
			phone,
			address
		FROM user_profiles
		WHERE user_id=$1
		`,
		userID,
	)

	var profile model.UserProfile

	err := row.Scan(
		&profile.ID,
		&profile.UserID,
		&profile.Email,
		&profile.Phone,
		&profile.Address,
	)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *ProfileRepository) Update(
	profile model.UserProfile,
) error {

	_, err := r.DB.Exec(
		`
		UPDATE user_profiles
		SET
			email=$1,
			phone=$2,
			address=$3
		WHERE user_id=$4
		`,
		profile.Email,
		profile.Phone,
		profile.Address,
		profile.UserID,
	)

	return err
}
