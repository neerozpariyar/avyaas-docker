package gorm

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

/*
ForgotPassword is a repository method responsible for handling the forgot password process, including
the generation of a one-time password (OTP) and updating the user's password in the database.

Parameters:
  - request: An instance of the ForgotPasswordRequest struct containing the necessary information
    for the forgot password process.

Returns:
  - error: An error indicating any issues encountered during the forgot password process.
    A nil error signifies a successful completion of the forgot password process.
*/
func (repo *Repository) ForgotPassword(request presenter.ForgotPasswordRequest) error {
	// Generate a new OTP string that will be used as a new password
	password, err := utils.GenerateOTP()
	if err != nil {
		return err
	}

	// Hash the password for secure storage
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// Update the user password with the new generated OTP's hash
	err = repo.db.Where("id = ?", request.UserID).Update("password", hashedPassword).Error
	if err != nil {
		return err
	}

	return nil
}
