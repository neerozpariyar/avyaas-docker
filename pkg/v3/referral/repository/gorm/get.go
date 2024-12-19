package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func (repo *Repository) GetReferralByID(id uint) (models.Referral, error) {
	var referral models.Referral

	// Retrieve the referral from the database based on given id
	err := repo.db.Where("id = ?", id).First(&referral).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Referral{}, fmt.Errorf("referral with referral id: '%d' not found", id)
		}

		return models.Referral{}, err
	}

	return referral, nil
}

func (repo *Repository) GetReferralByCode(code string) (models.Referral, error) {
	var referral models.Referral

	// Retrieve the referral from the database based on given id
	err := repo.db.Where("code = ?", code).First(&referral).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Referral{}, fmt.Errorf("referral with code: '%s' not found", code)
		}

		return models.Referral{}, err
	}

	uReferral, err := repo.CheckPendingReferralInTransaction(referral.ID)
	if err != nil {
		return models.Referral{}, err
	}

	return uReferral, nil
}

func (repo *Repository) CheckUserReferral(userID, referralID uint) (*models.UserReferral, error) {
	var userReferral *models.UserReferral

	err := repo.db.Where("user_id = ? AND referral_id = ?", userID, referralID).First(&userReferral).Error
	if err != nil {
		return nil, err
	}

	return userReferral, err
}

func (repo *Repository) CheckPendingReferralInTransaction(id uint) (models.Referral, error) {
	var err error
	var referralInTransactions []models.ReferralInTransaction
	var ritCount int64
	var referral models.Referral

	err = repo.db.Where("id = ?", id).First(&referral).Error
	if err != nil {
		return models.Referral{}, err
	}

	err = repo.db.Where("referral_id = ?", id).Find(&referralInTransactions).Count(&ritCount).Error
	if err != nil {
		return models.Referral{}, err
	}

	if ritCount == 0 {
		return referral, nil
	}

	for _, rit := range referralInTransactions {
		if rit.HoldTime.Unix() < time.Now().Unix() {
			err = repo.db.Unscoped().Where("id = ?", rit.ID).Delete(&models.ReferralInTransaction{}).Error
			if err != nil {
				return models.Referral{}, err
			}

			err = repo.db.Unscoped().Where("user_id = ? AND referral_id = ?", rit.UserID, rit.ReferralID).Delete(&models.UserReferral{}).Error
			if err != nil {
				return models.Referral{}, err
			}

			err = repo.db.Model(&models.Referral{}).Where("id = ?", id).Update("limit", referral.Limit+1).Error
			if err != nil {
				return models.Referral{}, err
			}

			referral, err = repo.GetReferralByID(id)
			if err != nil {
				return models.Referral{}, err
			}
		}
	}

	return referral, nil
}
