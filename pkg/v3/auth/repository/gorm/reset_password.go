package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"gorm.io/gorm"
)

func (repo *Repository) ResetPassword(request presenter.ResetPasswordRequest) error {
	// Hash the password for secure storage
	hashedPassword, err := utils.HashPassword(request.NewPassword)
	if err != nil {
		return err
	}

	var baseQuery *gorm.DB

	if isValidEmail := utils.IsValidEmail(request.Identity); isValidEmail {
		baseQuery = repo.db.Model(&models.User{}).Where("email = ?", request.Identity)
	} else if ok := utils.ContainsOnlyNumber(request.Identity); ok {
		baseQuery = repo.db.Model(&models.User{}).Where("phone = ?", request.Identity)
	}

	// Update the user password with the new generated OTP's hash
	err = baseQuery.Update("password", hashedPassword).Error
	if err != nil {
		return err
	}

	return nil
}
