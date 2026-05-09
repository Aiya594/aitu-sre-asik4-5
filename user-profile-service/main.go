package main

import (
	"os"

	"github.com/Aiya594/aitu-sre-asik4-5-user-profile/internal/config"
	"github.com/Aiya594/aitu-sre-asik4-5-user-profile/internal/handler"
	repository "github.com/Aiya594/aitu-sre-asik4-5-user-profile/internal/repositoy"
	"github.com/Aiya594/aitu-sre-asik4-5-user-profile/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	_ = godotenv.Load(".env")

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8085"
	}

	database, err := config.NewDB()

	if err != nil {
		panic(err)
	}

	repo := &repository.ProfileRepository{
		DB: database,
	}

	svc := &service.ProfileService{
		Repo: repo,
	}

	h := &handler.ProfileHandler{
		Service: svc,
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "profile ok",
		})
	})

	r.POST("/profile", h.Create)

	r.GET("/profile/:userID", h.Get)

	r.PUT("/profile/:userID", h.Update)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":" + port)
}
