package service

import (
	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/model"
	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/repository"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func (s *OrderService) CreateOrder(userID int, product string, amount float64) error {
	order := model.Order{
		UserID:  userID,
		Product: product,
		Amount:  amount,
	}

	return s.Repo.Create(order)
}

func (s *OrderService) GetOrders(userID int) ([]model.Order, error) {
	return s.Repo.GetByUser(userID)
}
