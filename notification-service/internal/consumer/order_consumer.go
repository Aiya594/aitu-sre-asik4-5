package consumer

import (
	"encoding/json"
	"log"

	"github.com/Aiya594/aitu-sre-asik4-5-notification/internal/service"
	"github.com/rabbitmq/amqp091-go"
)

type OrderCreatedEvent struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type OrderConsumer struct {
	Service *service.NotificationService
}

func (c *OrderConsumer) Consume(
	msgs <-chan amqp091.Delivery,
) {

	for msg := range msgs {

		var event OrderCreatedEvent

		err := json.Unmarshal(
			msg.Body,
			&event,
		)

		if err != nil {

			log.Printf(
				"[Consumer][ERROR] invalid message: %v",
				err,
			)

			continue
		}

		log.Printf(
			"[Consumer] received event for user=%d",
			event.UserID,
		)

		err = c.Service.Send(
			event.UserID,
			event.Message,
			event.Type,
		)

		if err != nil {

			log.Printf(
				"[Consumer][ERROR] notification failed: %v",
				err,
			)

			continue
		}

		log.Printf(
			"[Consumer] notification processed",
		)
	}
}
