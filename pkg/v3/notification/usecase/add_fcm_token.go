package usecase

import (
	"avyaas/internal/domain/models"
	"errors"
	"log"
)

func (uCase *usecase) AddFCMToken(fcmToken models.FCMToken) error {
	existingFCM, err := uCase.repo.GetFCMByUserIDAndToken(fcmToken.UserID, fcmToken.Token)
	if err != nil {
		log.Println("Error getting fcm token:", err)
	}
	if existingFCM != (models.FCMToken{}) {
		return errors.New("FCMToken for this user and token already exists")
	}
	err = uCase.repo.AddFCMToken(fcmToken)
	if err != nil {
		return err
	}
	err = uCase.FCMRegister(fcmToken)
	if err != nil {
		return err
	}
	return nil
}

func (uCase *usecase) GetFCMByUserIDAndToken(userID uint, token string) (fcmToken models.FCMToken, err error) {
	if _, err := uCase.repo.GetFCMByUserIDAndToken(userID, token); err != nil {
		return fcmToken, err
	}

	return fcmToken, nil
}
