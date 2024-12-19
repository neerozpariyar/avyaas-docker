package gorm

import (
	"avyaas/internal/domain/models"

	"time"

	"github.com/spf13/viper"
)

func (repo *Repository) CreateReferralInTransaction(userID, referralID, packageID uint) error {
	transaction := repo.db.Begin()
	holdTime := time.Now().Add(time.Minute * time.Duration(viper.GetInt("referralHoldTime")))

	err := transaction.Create(&models.ReferralInTransaction{
		UserID:     userID,
		PackageID:  packageID,
		ReferralID: referralID,
		HoldTime:   &holdTime,
	}).Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	referral, err := repo.GetReferralByID(referralID)
	if err != nil {
		transaction.Rollback()
		return err
	}

	err = transaction.Create(&models.UserReferral{
		UserID:     userID,
		ReferralID: referral.ID,
	}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	err = transaction.Model(&models.Referral{}).Where("id = ?", referralID).Update("limit", referral.Limit-1).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
