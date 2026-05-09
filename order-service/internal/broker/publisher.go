package broker

import (
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

type NotificationEvent struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type Publisher struct {
	Channel *amqp091.Channel
}

func (p *Publisher) PublishOrderCreated(
	event NotificationEvent,
) error {

	body, err := json.Marshal(event)

	if err != nil {
		return err
	}

	return p.Channel.Publish(
		"",
		"order_created",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
