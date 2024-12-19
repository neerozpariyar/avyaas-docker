package main

import (
	"avyaas/internal/config"
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	_notificationRepo "avyaas/pkg/v3/notification/repository/gorm"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

// to run websocket server do "make ws"
// ws://localhost:8081/notification/
func main() {
	app := fiber.New()

	app.Get("/notification", websocket.New(Consume))

	go func() {
		if err := app.Listen(":8081"); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("Server terminated")
}

func Consume(c *websocket.Conn) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Printf("failed to connect to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("failed to open a channel: %v", err)
		return
	}
	defer ch.Close()

	// Declare the queues for the consumers
	queues := map[string]string{
		"verified":   "verified-users-queue",
		"unverified": "unverified-users-queue",
		"course":     "course-users-queue",
	}

	for recipient, queueName := range queues {
		q, err := ch.QueueDeclare(
			queueName, // name
			true,      // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		if err != nil {
			log.Printf("failed to declare queue %s: %v", queueName, err)
			continue
		}

		// Consume messages from the queue
		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		if err != nil {
			log.Printf("failed to register a consumer for queue %s: %v", queueName, err)
			continue
		}

		// Handle messages in a separate goroutine
		go func(recipient string, msgs <-chan amqp.Delivery) {
			db := config.InitDB(true, false)
			dbInit := _notificationRepo.New(db)

			for msg := range msgs {
				notification := models.Notification{}
				err := json.Unmarshal(msg.Body, &notification)
				if err != nil {
					log.Printf("failed to unmarshal message for recipient %s: %v", recipient, err)
					continue
				}
				// Set consumed to true
				notification.Consumed = true

				// Update the notification in the database
				// db := *gorm.DB
				// Initialize the database

				// config.ConfigureViper()

				// Prepare the response
				response := presenter.ConsumerResponse{
					Recipient: recipient,
					Data:      []models.Notification{notification},
				}

				// Encode the response to JSON
				responseData, err := json.Marshal(response)
				if err != nil {
					log.Printf("failed to encode response to JSON: %v", err)
					continue
				}

				// Write the JSON response to the WebSocket connection
				if err := c.WriteMessage(websocket.TextMessage, responseData); err != nil {
					log.Printf("failed to write message to WebSocket: %v", err)
					continue
				}
				transaction := db.Begin()

				err = dbInit.UpdateNotification(notification)
				if err != nil {
					transaction.Rollback()
					log.Printf("failed to update notification: %v", err)
					// continue
				} else {
					transaction.Commit()
				}
			}
		}(recipient, msgs)

		// Wait for the WebSocket connection to close
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				log.Printf("WebSocket connection closed: %v", err)
				break
			}
		}
	}
}
