package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

/*
ChangePassword is a repository method responsible for updating a user's password in the underlying
database.

Parameters:
  - request: An instance of the ChangePasswordRequest struct containing the UserID and NewPassword
    details.

Returns:
  - error: An error indicating any issues encountered during the password update process.
    A nil error signifies a successful password update in the database.
*/
func (repo *Repository) ChangePassword(request presenter.ChangePasswordRequest) error {
	// Hash the password for secure storage
	password, err := utils.HashPassword(request.NewPassword)
	if err != nil {
		return err
	}

	// Initiate the password update with new hashed password
	err = repo.db.Model(&models.User{}).Where("id = ?", request.UserID).Update("password", password).Error
	if err != nil {
		return err
	}

	return nil
}
