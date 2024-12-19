package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

/*
SaveUserOTP saves the generated OTP (One-Time Password) for a given phone number in the database.
The function creates a new UserOtp record with the provided phone number, OTP, and expiration time.

Parameters:
  - phone: The phone number associated with the OTP.
  - otp: The generated One-Time Password.

Returns:
  - err: An error if any operation fails during the creation of the OTP record.
*/
func (repo *Repository) SaveUserOTP(identity, otp string) error {
	userOTP, err := repo.CheckUserOTP(identity)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if userOTP != nil {
		/* Update the existing OTP with the new one, an instance with otp for phone number already
		exists */
		return repo.db.Model(&models.UserOtp{}).Where("identity = ?", identity).Updates(&models.UserOtp{
			Timestamp: models.Timestamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			OTP:       otp,
			ExpiresAt: time.Now().Add(time.Minute * 5),
		}).Error
	}

	return repo.db.Create(&models.UserOtp{
		Identity:  identity,
		OTP:       otp,
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}).Error
}

/*
CheckUserOTP checks and retrieves the UserOtp record associated with the provided phone number from
the database.

Parameters:
  - phone: A phone number string

Returns:
  - otp: A pointer to an models.UserOtp containing the OTP information.
  - err: An error, if any occurred during the process.
*/
func (repo *Repository) CheckUserOTP(identity string) (*models.UserOtp, error) {
	var userOTP *models.UserOtp

	err := repo.db.Model(&models.UserOtp{}).Where("identity = ?", identity).First(&userOTP).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return userOTP, nil
}
