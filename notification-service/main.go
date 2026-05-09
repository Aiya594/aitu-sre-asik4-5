package main

import (
	"os"

	"github.com/Aiya594/aitu-sre-asik4-5-notification/internal/broker"
	"github.com/Aiya594/aitu-sre-asik4-5-notification/internal/config"
	"github.com/Aiya594/aitu-sre-asik4-5-notification/internal/consumer"
	"github.com/Aiya594/aitu-sre-asik4-5-notification/internal/repository"
	"github.com/Aiya594/aitu-sre-asik4-5-notification/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load(".env")

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8084"
	}

	db, err := config.NewDB()

	if err != nil {
		panic(err)
	}

	repo := &repository.NotificationRepository{
		DB: db,
	}

	svc := &service.NotificationService{
		Repo: repo,
	}

	/*
		RabbitMQ
	*/

	rabbitURL := os.Getenv("RABBITMQ_URL")

	conn, ch, err := broker.ConnectRabbitMQ(
		rabbitURL,
	)

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	defer ch.Close()

	/*
		DECLARE QUEUE
	*/

	q, err := ch.QueueDeclare(
		"order_created",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	orderConsumer := &consumer.OrderConsumer{
		Service: svc,
	}

	go orderConsumer.Consume(msgs)

	/*
		HTTP SERVER
	*/

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "notification ok",
		})
	})

	r.Run(":" + port)
}
