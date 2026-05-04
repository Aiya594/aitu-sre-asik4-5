package service

import (
	"log"

	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/model"
	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/repository"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func (s *OrderService) CreateOrder(userID int, product string, amount float64) error {
	log.Printf("CreateOrder called: userID=%d product=%s amount=%.2f", userID, product, amount)

	order := model.Order{
		UserID:  userID,
		Product: product,
		Amount:  amount,
	}

	err := s.Repo.Create(order)
	if err != nil {
		log.Printf("ERROR creating order: userID=%d product=%s amount=%.2f err=%v",
			userID, product, amount, err)
		return err
	}

	log.Printf("Order created successfully: userID=%d product=%s amount=%.2f",
		userID, product, amount)

	return nil
}

func (s *OrderService) GetOrders(userID int) ([]model.Order, error) {
	log.Printf("GetOrders called: userID=%d", userID)

	orders, err := s.Repo.GetByUser(userID)
	if err != nil {
		log.Printf("ERROR fetching orders: userID=%d err=%v", userID, err)
		return nil, err
	}

	log.Printf("GetOrders success: userID=%d count=%d", userID, len(orders))

	return orders, nil
}
