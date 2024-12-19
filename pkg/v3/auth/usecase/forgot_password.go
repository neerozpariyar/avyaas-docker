package usecase

import (
	"avyaas/internal/domain/presenter"
)

/*
ForgotPassword is a use case function responsible for initiating the forgot password process for a
user.

Parameters:
  - request: An instance of the ForgotPasswordRequest struct containing the necessary information
    for initiating the forgot password process, including the user's phone number.

Returns:s
  - error: An error indicating any issues encountered during the initiation of the forgot password
    process. A nil error signifies a successful initiation of the forgot password process.
*/
func (uCase *usecase) ForgotPassword(request presenter.ForgotPasswordRequest) error {
	// Check and retrieve if the user with given phone number exists
	user, err := uCase.accountRepo.GetUserByPhone(request.Phone)
	if err != nil {
		return err
	}

	// Set the request userID with the retrieved user's ID
	request.UserID = user.ID
	err = uCase.repository.ForgotPassword(request)
	if err != nil {
		return err
	}

	return err
}
