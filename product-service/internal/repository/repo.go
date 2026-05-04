package repository

import (
	"database/sql"

	"github.com/Aiya594/aitu-sre-asik4-5-product/internal/model"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) Create(p model.Product) error {
	_, err := r.DB.Exec(
		"INSERT INTO products (name, price, stock) VALUES ($1, $2, $3)",
		p.Name, p.Price, p.Stock,
	)
	return err
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var p model.Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
	row := r.DB.QueryRow(
		"SELECT id, name, price, stock FROM products WHERE id=$1",
		id,
	)

	var p model.Product
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *ProductRepository) UpdateStock(id int, stock int) error {
	_, err := r.DB.Exec(
		"UPDATE products SET stock=$1 WHERE id=$2",
		stock, id,
	)
	return err
}
