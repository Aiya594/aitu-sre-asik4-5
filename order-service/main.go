package main

import (
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

	//port := os.Getenv("APP_PORT")
	database, err := configs.NewDB()
	if err != nil {
		panic(err)
	}

	repo := &repository.OrderRepository{DB: database}
	svc := &service.OrderService{Repo: repo}
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

	r.Run(":8081")
}
