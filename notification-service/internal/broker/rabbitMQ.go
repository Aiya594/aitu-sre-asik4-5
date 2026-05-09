package broker

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ(
	url string,
) (*amqp091.Connection, *amqp091.Channel, error) {

	conn, err := amqp091.Dial(url)

	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, nil, err
	}

	log.Printf("[RabbitMQ] connected")

	return conn, ch, nil
}
