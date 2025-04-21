package config

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

func LoadEnv() {
	_ = os.Setenv("RABBITMQ_URI", os.Getenv("RABBITMQ_URI"))
}

func ConnectRabbitMQ() {
	uri := os.Getenv("RABBITMQ_URI")
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	RabbitConn = conn

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	RabbitChannel = ch

	log.Println("Connected to RabbitMQ")
}
