package main

import (
	"log"

	global "github.com/fiqrioemry/microservice-ecommerce/server/notification-service/internal/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/notification-service/internal/consumer"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"
)

func main() {
	config.LoadEnv()
	global.ConnectRabbitMQ()
	config.InitMailer()
	log.Println("Starting Notification Service...")

	consumer.ConsumeOrderCreated()

	select {} // biar gak exit
}
