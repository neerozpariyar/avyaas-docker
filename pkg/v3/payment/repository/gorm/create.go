package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) CreatePayment(request *models.Payment) error {
	// var existingPayment models.Payment
	// merchant := strings.ToUpper(request.MerchantType)
	// if merchant == "KHALTI" {
	// 	// If merchant type is khalti, check for uuid only
	// 	result := repo.db.Where("transaction_uuid = ?", request.TransactionUUID).First(&existingPayment)
	// 	if result.Error == nil {
	// 		// If a record is found, return an error
	// 		return errors.New("payment already exists")
	// 	}
	// } else if merchant == "ESEWA" {
	// 	// If merchant type is esewa, check all three
	// 	result := repo.db.Where("merchant_type = ? AND transaction_id = ? AND transaction_uuid = ?", request.MerchantType, request.TransactionID, request.TransactionUUID).First(&existingPayment)
	// 	if result.Error == nil {
	// 		// If a record is found, return an error
	// 		return errors.New("payment already exists")
	// 	}
	// }

	// If no matching record is found, create the payment
	return repo.db.Create(&request).Error
}
