package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func (uc *usecase) SendFCMNotification(notification models.Notification) error {
	// notification, err := uc.repo.GetNotificationByID(notificationID)
	// if err != nil {
	// 	return err
	// }
	var jsonStr = map[string]string{"app": viper.GetString("notificationService.app"), "title": notification.Title, "body": notification.Description}

	body, err := json.Marshal(jsonStr)
	if err != nil {
		return err
	}

	token, err := utils.GenerateFCMJWT()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", viper.GetString("notificationService.notification_url"), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fiber.ErrBadRequest
	}

	return nil
}
