package handler

import (
	"strconv"

	"github.com/Aiya594/aitu-sre-asik4-5-user-profile/internal/service"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	Service *service.ProfileService
}

func (h *ProfileHandler) Create(c *gin.Context) {

	var req struct {
		UserID  int    `json:"user_id"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}

	if err := c.BindJSON(&req); err != nil {

		c.JSON(400, gin.H{
			"error": "invalid request",
		})

		return
	}

	err := h.Service.CreateProfile(
		req.UserID,
		req.Email,
		req.Phone,
		req.Address,
	)

	if err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"status": "profile created",
	})
}

func (h *ProfileHandler) Get(c *gin.Context) {

	userID, _ := strconv.Atoi(
		c.Param("userID"),
	)

	profile, err := h.Service.GetProfile(userID)

	if err != nil {

		c.JSON(404, gin.H{
			"error": "profile not found",
		})

		return
	}

	c.JSON(200, profile)
}

func (h *ProfileHandler) Update(c *gin.Context) {

	userID, _ := strconv.Atoi(
		c.Param("userID"),
	)

	var req struct {
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}

	if err := c.BindJSON(&req); err != nil {

		c.JSON(400, gin.H{
			"error": "invalid request",
		})

		return
	}

	err := h.Service.UpdateProfile(
		userID,
		req.Email,
		req.Phone,
		req.Address,
	)

	if err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"status": "profile updated",
	})
}
