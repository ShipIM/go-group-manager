package rabbit

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbit(uri string, exchange string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate connection: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	return conn, nil
}
