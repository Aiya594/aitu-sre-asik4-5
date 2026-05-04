package main

import (
	"os"

	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/configs"
	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/handler"
	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/repository"
	"github.com/Aiya594/aitu-sre-asik4-5-auth/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	_ = godotenv.Load(".env")

	port := os.Getenv("APP_PORT")

	database, err := configs.NewDB()
	if err != nil {
		panic(err)
	}

	repo := &repository.UserRepository{DB: database}
	svc := &service.AuthService{Repo: repo}
	h := &handler.AuthHandler{Service: svc}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "auth ok"})
	})

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":" + port)
}
