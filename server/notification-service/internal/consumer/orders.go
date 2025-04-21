package consumer

import (
	"encoding/json"
	"log"

	"github.com/fiqrioemry/microservice-ecommerce/server/notification-service/internal/config"
)

type OrderCreatedPayload struct {
	OrderID   string  `json:"orderId"`
	UserID    string  `json:"userId"`
	Total     float64 `json:"total"`
	CreatedAt string  `json:"createdAt"`
}

func ConsumeOrderCreated() {
	q, err := config.RabbitChannel.QueueDeclare(
		"order.created", // queue name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	msgs, err := config.RabbitChannel.Consume(
		q.Name,
		"notification-service", // consumer
		true,                   // auto-ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	log.Println("Listening to order.created events...")

	go func() {
		for d := range msgs {
			var payload OrderCreatedPayload
			if err := json.Unmarshal(d.Body, &payload); err != nil {
				log.Println("Invalid payload:", err)
				continue
			}

			log.Printf("📩 New order.created: OrderID=%s UserID=%s Total=%.2f\n", payload.OrderID, payload.UserID, payload.Total)
			// TODO: simpan ke DB, kirim email, dll
		}
	}()
}
