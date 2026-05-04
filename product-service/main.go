package main

import (
	"os"

	"github.com/Aiya594/aitu-sre-asik4-5-product/internal/config"
	"github.com/Aiya594/aitu-sre-asik4-5-product/internal/handler"
	"github.com/Aiya594/aitu-sre-asik4-5-product/internal/repository"
	"github.com/Aiya594/aitu-sre-asik4-5-product/internal/service"
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
		port = "8082"
	}

	database, err := config.NewDB()
	if err != nil {
		panic(err)
	}

	repo := &repository.ProductRepository{DB: database}
	svc := &service.ProductService{Repo: repo}
	h := &handler.ProductHandler{Service: svc}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "product ok"})
	})

	r.POST("/product", h.Create)
	r.GET("/products", h.GetAll)
	r.GET("/products/:id", h.GetByID)
	r.PUT("/products/:id/decrease-stock", h.DecreaseStock)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":" + port)
}
