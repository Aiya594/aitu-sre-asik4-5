package broker

import (
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ(
	url string,
) (*amqp091.Connection, *amqp091.Channel, error) {

	var conn *amqp091.Connection
	var ch *amqp091.Channel
	var err error

	for i := 0; i < 10; i++ {

		log.Printf(
			"[RabbitMQ] connecting attempt %d...",
			i+1,
		)

		conn, err = amqp091.Dial(url)

		if err == nil {

			ch, err = conn.Channel()

			if err == nil {

				log.Printf(
					"[RabbitMQ] connected",
				)

				return conn, ch, nil
			}
		}

		log.Printf(
			"[RabbitMQ] not ready, retrying...",
		)

		time.Sleep(5 * time.Second)
	}

	return nil, nil, err
}
