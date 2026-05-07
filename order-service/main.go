package main

import (
	"os"

	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/client"
	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/configs"
	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/handler"
	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/repository"
	"github.com/Aiya594/aitu-sre-asik4-5-order/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	_ = godotenv.Load(".env")

	// if err != nil {
	// 	panic(err)
	// }
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8081"
	}

	productUrl := os.Getenv("PRODUCT_SERVICE_URL")
	paymentURL := os.Getenv("PAYMENT_SERVICE_URL")

	database, err := configs.NewDB()
	if err != nil {
		panic(err)
	}

	repo := &repository.OrderRepository{DB: database}

	product := client.NewProductClient(productUrl)
	payment := client.NewPaymentClient(paymentURL)

	svc := &service.OrderService{
		Repo:    repo,
		Product: product,
		Payment: payment,
	}

	h := &handler.OrderHandler{Service: svc}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "order ok"})
	})

	auth := r.Group("/")
	auth.Use(configs.AuthMiddleware())

	auth.POST("/order", h.Create)
	auth.GET("/orders", h.GetOrders)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":" + port)
}
