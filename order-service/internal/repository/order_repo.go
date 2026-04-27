package repository

import (
	"database/sql"

	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/model"
)

type OrderRepository struct {
	DB *sql.DB
}

func (r *OrderRepository) Create(order model.Order) error {
	_, err := r.DB.Exec(
		"INSERT INTO orders (user_id, product, amount) VALUES ($1, $2, $3)",
		order.UserID, order.Product, order.Amount,
	)
	return err
}

func (r *OrderRepository) GetByUser(userID int) ([]model.Order, error) {
	rows, err := r.DB.Query(
		"SELECT id, user_id, product, amount FROM orders WHERE user_id=$1",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order

	for rows.Next() {
		var o model.Order
		rows.Scan(&o.ID, &o.UserID, &o.Product, &o.Amount)
		orders = append(orders, o)
	}

	return orders, nil
}
