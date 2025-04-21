package main

import (
	"log"

	"github.com/fiqrioemry/microservice-ecommerce/server/notification-service/internal/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/notification-service/internal/consumer"
)

func main() {
	config.LoadEnv()
	config.ConnectRabbitMQ()

	log.Println("Starting Notification Service...")

	consumer.ConsumeOrderCreated()

	select {} // biar gak exit
}
