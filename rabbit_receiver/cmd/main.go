package main

import (
	"log"

	"github.com/ShipIM/go-group-manager/rabbit_receiver/internal/rabbit"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("can not initialize configs: %s", err.Error())
	}

	var (
		rabbitAddress   = viper.GetString("rabbit.address")
		rabbitExchanger = viper.GetString("rabbit.exchanger")
		rabbitInQueue   = viper.GetString("rabbit.in_queue")
	)

	connMq, err := rabbit.InitRabbit(rabbitAddress, rabbitExchanger, rabbitInQueue)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %s", err)
	}
	defer connMq.Close()

	chMq, err := connMq.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %s", err)
	}
	defer chMq.Close()

	msgs, err := chMq.Consume(
		rabbitInQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to consume messages: %s", err)
	}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for students. To exit press CTRL+C")

	var forever chan struct{}
	<-forever
}

func initConfig() error {
	viper.AddConfigPath("../config")
	viper.SetConfigName("application")

	return viper.ReadInConfig()
}
