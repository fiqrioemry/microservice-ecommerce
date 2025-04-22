package consumer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fiqrioemry/microservice-ecommerce/server/notification-service/internal/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
)

type OrderCreatedPayload struct {
	OrderID   string  `json:"orderId"`
	UserID    string  `json:"userId"`
	Total     float64 `json:"total"`
	CreatedAt string  `json:"createdAt"`
	// Optional: UserEmail string `json:"userEmail"`
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

			log.Printf("📩 New order.created: OrderID=%s UserID=%s Total=%.2f", payload.OrderID, payload.UserID, payload.Total)

			// Sementara: kirim email ke alamat hardcoded
			email := "user@example.com" // ganti dengan payload.UserEmail jika tersedia
			message := fmt.Sprintf("Hi 👋, order #%s telah berhasil dibuat sebesar Rp%.0f.", payload.OrderID, payload.Total)
			err := utils.SendEmail(email, "", message)
			if err != nil {
				log.Println("❌ Gagal kirim email:", err)
			} else {
				log.Println("✅ Email notifikasi berhasil dikirim ke:", email)
			}
		}
	}()
}
