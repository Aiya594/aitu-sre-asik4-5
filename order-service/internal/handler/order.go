package handler

import (
	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Service *service.OrderService
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req struct {
		Product string  `json:"product"`
		Amount  float64 `json:"amount"`
	}

	c.BindJSON(&req)

	userID := c.GetInt("user_id")

	err := h.Service.CreateOrder(userID, req.Product, req.Amount)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "order created"})
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	userID := c.GetInt("user_id")

	orders, err := h.Service.GetOrders(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orders)
}
