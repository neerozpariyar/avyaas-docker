package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) AddFCMToken(fcmToken models.FCMToken) (err error) {
	fcmToken.Registered = true
	return repo.db.Create(&fcmToken).Error
}

func (repo *Repository) GetFCMByUserIDAndToken(userID uint, token string) (fcmToken models.FCMToken, err error) {
	err = repo.db.Where("user_id = ? AND token = ?", userID, token).First(&fcmToken).Error
	return fcmToken, err
}
