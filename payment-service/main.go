package main

import (
	"os"

	"github.com/Aiya594/aitu-sre-asik4-5-payment/internal/config"
	"github.com/Aiya594/aitu-sre-asik4-5-payment/internal/handler"
	"github.com/Aiya594/aitu-sre-asik4-5-payment/internal/repository"
	"github.com/Aiya594/aitu-sre-asik4-5-payment/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	_ = godotenv.Load(".env")
	// if err != nil {
	// 	panic(err)
	// }

	port := os.Getenv("PAYMENT_SERVICE_PORT")
	if port == "" {
		port = "8083"
	}

	database, err := config.NewDB()

	if err != nil {
		panic(err)
	}

	repo := &repository.PaymentRepository{
		DB: database,
	}

	svc := &service.PaymentService{
		Repo: repo,
	}

	h := &handler.PaymentHandler{
		Service: svc,
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "payment ok",
		})
	})

	r.POST("/payment", h.Process)

	r.GET("/payments", h.GetAll)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":" + port)
}
