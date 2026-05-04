package handler

import (
	"strconv"

	"github.com/Aiya594/aitu-sre-asik4-5-product/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *service.ProductService
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock int     `json:"stock"`
	}

	c.BindJSON(&req)

	err := h.Service.CreateProduct(req.Name, req.Price, req.Stock)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "product created"})
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.Service.GetProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.Service.GetProduct(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.JSON(200, product)
}

func (h *ProductHandler) DecreaseStock(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Amount int `json:"amount"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	product, err := h.Service.GetProduct(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "product not found"})
		return
	}

	if product.Stock < req.Amount {
		c.JSON(400, gin.H{"error": "not enough stock"})
		return
	}

	err = h.Service.DecreaseStock(id, req.Amount)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to update stock"})
		return
	}

	c.JSON(200, gin.H{"status": "stock updated"})
}
