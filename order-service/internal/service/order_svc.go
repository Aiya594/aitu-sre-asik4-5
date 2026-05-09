package service

import (
	"log"

	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/broker"
	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/client"
	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/model"
	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/repository"
	"github.com/google/uuid"
)

type OrderService struct {
	Repo      *repository.OrderRepository
	Product   client.ProductClient
	Payment   client.PaymentClient
	Publisher *broker.Publisher
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

func (s *OrderService) CreateOrder(userID int, productID int, productName string, amount float64) error {

	log.Printf("[OrderService] CreateOrder user=%d productID=%d amount=%.2f", userID, productID, amount)

	log.Printf("[UseCase] CreateOrder user=%d product=%d amount=%.2f", userID, productID, amount)

	orderID := uuid.New().String()
	err := s.Payment.ProcessPayment(
		orderID,
		userID,
		amount,
	)

	if err != nil {
		log.Printf(
			"[OrderService][ERROR] payment failed: %v",
			err,
		)

		return err
	}

	// 1. reserve stock FIRST (distributed transaction step)
	err = s.Product.DecreaseStock(productID, int(amount))
	if err != nil {
		log.Printf("[UseCase][ERROR] stock reservation failed: %v", err)
		return err
	}

	// 2. CREATE ORDER IN DB
	order := model.Order{
		ID:      orderID,
		UserID:  userID,
		Product: productName,
		Amount:  amount,
	}

	err = s.Repo.Create(order)
	if err != nil {
		log.Printf("[OrderService][ERROR] DB insert failed: %v", err)
		return err
	}

	err = s.Publisher.PublishOrderCreated(
		broker.NotificationEvent{
			UserID:  userID,
			Message: "Your order was created successfully",
			Type:    "email",
		},
	)

	if err != nil {

		log.Printf(
			"[OrderService][WARN] failed to publish notification",
		)
	}

	log.Printf("[OrderService] order created successfully for user=%d", userID)
	return nil
}
