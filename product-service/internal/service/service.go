package service

import (
	"errors"
	"log"

	"github.com/Aiya594/aitu-sre-asik4-5-product/internal/model"
	"github.com/Aiya594/aitu-sre-asik4-5-product/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

/*
Business rules:
- price must be > 0
- stock cannot be negative
- product name must not be empty
*/
func (s *ProductService) CreateProduct(name string, price float64, stock int) error {

	log.Printf("[ProductService] CreateProduct called: name=%s price=%.2f stock=%d", name, price, stock)

	// Validation layer (business logic)
	if name == "" {
		log.Printf("[ProductService][ERROR] empty product name")
		return errors.New("product name cannot be empty")
	}

	if price <= 0 {
		log.Printf("[ProductService][ERROR] invalid price: %.2f", price)
		return errors.New("price must be greater than 0")
	}

	if stock < 0 {
		log.Printf("[ProductService][ERROR] invalid stock: %d", stock)
		return errors.New("stock cannot be negative")
	}

	product := model.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}

	err := s.Repo.Create(product)
	if err != nil {
		log.Printf("[ProductService][ERROR] DB insert failed: %v", err)
		return err
	}

	log.Printf("[ProductService] product created successfully: %s", name)
	return nil
}

func (s *ProductService) GetProducts() ([]model.Product, error) {

	log.Printf("[ProductService] GetProducts called")

	products, err := s.Repo.GetAll()
	if err != nil {
		log.Printf("[ProductService][ERROR] failed to fetch products: %v", err)
		return nil, err
	}

	log.Printf("[ProductService] fetched %d products", len(products))
	return products, nil
}

func (s *ProductService) GetProduct(id int) (*model.Product, error) {

	log.Printf("[ProductService] GetProduct called: id=%d", id)

	if id <= 0 {
		log.Printf("[ProductService][ERROR] invalid product id: %d", id)
		return nil, errors.New("invalid product id")
	}

	product, err := s.Repo.GetByID(id)
	if err != nil {
		log.Printf("[ProductService][ERROR] product not found: id=%d err=%v", id, err)
		return nil, err
	}

	// business rule example: hide out-of-stock products (optional logic)
	if product.Stock == 0 {
		log.Printf("[ProductService][WARN] product out of stock: id=%d", id)
	}

	log.Printf("[ProductService] product retrieved: id=%d name=%s", id, product.Name)
	return product, nil
}
