package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

/*
ChangePassword is a use case function responsible for processing a user's request to change their
password. It performs several validation checks, including verifying the correctness of the current
password, ensuring that the new password and its confirmation match, and assessing the strength of
the new password.

Parameters:
  - request: An instance of the ChangePasswordRequest struct containing the necessary information
    for the password change operation, including UserID, CurrentPassword, NewPassword, and ConfirmPassword.

Returns:
  - map[string]string: A map of error messages where keys represent specific fields or operations,
    and values contain corresponding error details. An empty map indicates a successful password change.
*/
func (uCase *usecase) ChangePassword(request presenter.ChangePasswordRequest) map[string]string {
	errMap := make(map[string]string)

	// Check if the user with given ID exists
	user, err := uCase.repo.GetUserByID(request.UserID)
	if err != nil {
		errMap["user"] = err.Error()
		return errMap
	}

	// Check if the current password hash match with that in the database
	if match := utils.CheckPasswordHash(request.CurrentPassword, user.Password); !match {
		errMap["currentPassword"] = "old password doesn't match"
		return errMap
	}

	// Check if both new password matches
	if request.NewPassword != request.ConfirmPassword {
		errMap["newPassword"] = "passwords do not match"
		errMap["confirmPassword"] = "passwords do not match"
		return errMap
	}

	// Check the strength of the provided password
	err = utils.CheckPasswordStrength(request.NewPassword)
	if err != nil {
		errMap["password"] = err.Error()
		return errMap
	}

	// Invoke the change password repo
	err = uCase.repo.ChangePassword(request)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
