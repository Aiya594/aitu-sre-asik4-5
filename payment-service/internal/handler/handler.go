package handler

import (
	"github.com/Aiya594/aitu-sre-asik4-5-payment/internal/service"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	Service *service.PaymentService
}

func (h *PaymentHandler) Process(c *gin.Context) {

	var req struct {
		OrderID int     `json:"order_id"`
		UserID  int     `json:"user_id"`
		Amount  float64 `json:"amount"`
	}

	if err := c.BindJSON(&req); err != nil {

		c.JSON(400, gin.H{
			"error": "invalid request",
		})

		return
	}

	err := h.Service.ProcessPayment(
		req.OrderID,
		req.UserID,
		req.Amount,
	)

	if err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"status": "payment successful",
	})
}

func (h *PaymentHandler) GetAll(c *gin.Context) {

	payments, err := h.Service.GetPayments()

	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, payments)
}
