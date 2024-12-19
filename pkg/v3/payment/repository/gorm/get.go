package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetUserPaymentByCoursePackage(userID, courseID, packageID uint) (payment *models.Payment, err error) {
	err = repo.db.Where("user_id = ? AND course_id = ? AND package_id = ?", userID, courseID, packageID).First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("payment for package: '%d' of course: '%d' by user:'%d' not found", packageID, courseID, userID)
		}
		return nil, err
	}

	return payment, nil
}

func (repo *Repository) GetPaymentByTransactionID(transactionID string) (*models.Payment, error) {
	var payment *models.Payment

	err := repo.db.Where("transaction_id = ?", transactionID).First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("payment with invoice number: '%s'", transactionID)
		}
		return nil, err
	}

	return payment, nil
}
