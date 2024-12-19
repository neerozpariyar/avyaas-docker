package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func (uCase *usecase) FCMRegister(fcmToken models.FCMToken) error {
	tokens := [1]string{fcmToken.Token}
	jsonStr := map[string]interface{}{
		"app":    viper.GetString("notificationService.app"),
		"tokens": tokens,
		"topic":  "avyaas",
	}
	jsonValue, _ := json.Marshal(jsonStr)
	token, err := utils.GenerateFCMJWT()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", viper.GetString("notificationService.addFcmUrl"), bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Println("Error creating notification request", err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error adding fcm to server: ", err)
		return err
	}
	log.Println(response.StatusCode)
	return nil
}
