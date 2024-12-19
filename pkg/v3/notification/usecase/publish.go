package usecase

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

func (uCase *usecase) PublishNotification(notificationID uint) error {
	notification, err := uCase.repo.GetNotificationByID(notificationID)
	if err != nil {
		return fmt.Errorf("failed to fetch notification details: %v", err)
	}
	notification.NotificationType = strings.ToLower(notification.NotificationType)
	if notification.NotificationType != "announcement" && notification.NotificationType != "push notification" {
		return nil
	}
	if notification.NotificationType == "push notification" {
		uCase.SendFCMNotification(notification)

	}
	if notification.NotificationType == "announcement" {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			return fmt.Errorf("failed to open a channel: %v", err)
		}
		defer ch.Close()

		args := make(amqp.Table)
		args["x-delayed-type"] = "direct" // Behavior type, could be direct, topic, fanout, etc.

		err = ch.ExchangeDeclare(
			"notification",      // name
			"x-delayed-message", // type
			true,                // durable
			false,               // auto-deleted
			false,               // internal
			false,               // no-wait
			args,                // arguments
		)
		if err != nil {
			return fmt.Errorf("failed to declare an exchange: %v", err)
		}

		body, err := json.Marshal(notification)
		if err != nil {
			return fmt.Errorf("failed to marshal notification: %v", err)
		}

		var routingKey string

		switch notification.Recipient {
		case "verified":
			routingKey = "verified-users"
			verifiedUsers, err := uCase.repo.GetVerifiedUsers()
			if err != nil {
				return fmt.Errorf("failed to fetch verified users: %v", err)
			}
			// Publish notifications to RabbitMQ queue for verified users
			for _, user := range verifiedUsers {
				if err := publishNotificationToUser(ch, routingKey, body, user.ID, *notification.ScheduledDate); err != nil {
					return err
				}
			}
		case "unverified":
			routingKey = "unverified-users"
			// Fetch unverified users from the database
			unverifiedUsers, err := uCase.repo.GetUnverifiedUsers()
			if err != nil {
				return fmt.Errorf("failed to fetch unverified users: %v", err)
			}
			// Publish notifications to RabbitMQ queue for unverified users
			for _, user := range unverifiedUsers {
				if err := publishNotificationToUser(ch, routingKey, body, user.ID, *notification.ScheduledDate); err != nil {
					return err
				}
			}
		case "course":
			routingKey = "course-users"
			if notification.CourseID == 0 {
				return fmt.Errorf("course ID not provided for course type notification")
			}
			// Fetch users subscribed to the specified course ID from the database
			courseUsers, err := uCase.repo.GetUsersByCourseID(notification.CourseID)
			if err != nil {
				return fmt.Errorf("failed to fetch users subscribed to course: %v", err)
			}
			// Publish notifications to RabbitMQ queue for course users
			for _, user := range courseUsers {
				if err := publishNotificationToUser(ch, routingKey, body, user.ID, *notification.ScheduledDate); err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("invalid recipient type: %s", notification.Recipient)
		}

		// go uCase.Consume()

		return err
	}
	return err
}
func publishNotificationToUser(ch *amqp.Channel, routingKey string, body []byte, userID uint, scheduledDate time.Time) error {

	headers := make(amqp.Table)
	if !scheduledDate.IsZero() && time.Now().Before(scheduledDate) {
		// Calculate the delay in milliseconds
		delay := time.Until(scheduledDate).Milliseconds()
		headers["x-delay"] = delay
	}
	queueName := fmt.Sprintf("%s-queue", routingKey)

	_, err := ch.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	// Bind the queue to the exchange with the routing key
	err = ch.QueueBind(
		queueName,      // queue name
		routingKey,     // routing key
		"notification", // exchange name
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind the queue: %v", err)
	}

	// Publish the message with the determined routing key
	err = ch.Publish(
		"notification", // exchange
		routingKey,     // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Headers:     headers,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}
	return err
}
