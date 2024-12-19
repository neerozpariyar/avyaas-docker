package usecase

import "avyaas/utils"

/*
ResendUserOTP checks if a user with the provided phone number exists, and if so, triggers the
process to resend a new OTP for verification.

Parameters:
  - phone: The user's phone number string.
  - err: An error occured during the process, if any.
*/
func (uCase *usecase) ResendUserOTP(identity string) error {
	// Check if the user with given phone number exists
	// _, err := uCase.accountRepo.GetUserByPhone(phone)
	// if err != nil {
	// 	return err
	// }

	if isValidEmail := utils.IsValidEmail(identity); isValidEmail {
		_, err := uCase.accountRepo.GetUserByEmail(identity)
		if err != nil {
			return err
		}
	} else if ok := utils.ContainsOnlyNumber(identity); ok {
		_, err := uCase.accountRepo.GetUserByPhone(identity)
		if err != nil {
			return err
		}
	}

	// Invoke otp resend repo
	return uCase.repository.ResendUserOTP(identity)
}
