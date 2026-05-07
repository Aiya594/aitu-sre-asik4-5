package repository

import (
	"database/sql"

	"github.com/Aiya594/aitu-sre-asik4-5-payment/internal/model"
)

type PaymentRepository struct {
	DB *sql.DB
}

func (r *PaymentRepository) Create(payment model.Payment) error {

	_, err := r.DB.Exec(
		`
		INSERT INTO payments (
			order_id,
			user_id,
			amount,
			status
		)
		VALUES ($1, $2, $3, $4)
		`,
		payment.OrderID,
		payment.UserID,
		payment.Amount,
		payment.Status,
	)

	return err
}

func (r *PaymentRepository) GetAll() ([]model.Payment, error) {

	rows, err := r.DB.Query(
		`
		SELECT
			id,
			order_id,
			user_id,
			amount,
			status
		FROM payments
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var payments []model.Payment

	for rows.Next() {

		var p model.Payment

		rows.Scan(
			&p.ID,
			&p.OrderID,
			&p.UserID,
			&p.Amount,
			&p.Status,
		)

		payments = append(payments, p)
	}

	return payments, nil
}
