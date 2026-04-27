package handler

import (
	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.BindJSON(&req)

	err := h.Service.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "user created"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.BindJSON(&req)

	token, err := h.Service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
