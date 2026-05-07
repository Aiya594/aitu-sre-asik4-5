package service

import (
	"errors"
	"log"

	"github.com/Aiya594/aitu-sre-asik4-5-payment/internal/model"
	"github.com/Aiya594/aitu-sre-asik4-5-payment/internal/repository"
)

type PaymentService struct {
	Repo *repository.PaymentRepository
}

func (s *PaymentService) ProcessPayment(
	orderID int,
	userID int,
	amount float64,
) error {

	log.Printf(
		"[PaymentService] ProcessPayment order=%d user=%d amount=%.2f",
		orderID,
		userID,
		amount,
	)

	if orderID <= 0 {
		return errors.New("invalid order id")
	}

	if userID <= 0 {
		return errors.New("invalid user id")
	}

	if amount <= 0 {
		return errors.New("invalid amount")
	}

	/*
		Payment simulation:
		- if amount > 10000 -> reject
		- else approve
	*/

	status := "paid"

	if amount > 10000 {
		status = "rejected"
	}

	payment := model.Payment{
		OrderID: orderID,
		UserID:  userID,
		Amount:  amount,
		Status:  status,
	}

	err := s.Repo.Create(payment)
	if err != nil {
		log.Printf("[PaymentService][ERROR] DB insert failed: %v", err)
		return err
	}

	if status == "rejected" {
		log.Printf("[PaymentService] payment rejected")
		return errors.New("payment rejected")
	}

	log.Printf("[PaymentService] payment successful")

	return nil
}

func (s *PaymentService) GetPayments() ([]model.Payment, error) {

	return s.Repo.GetAll()
}
