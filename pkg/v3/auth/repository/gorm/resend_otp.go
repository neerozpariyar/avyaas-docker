package gorm

import (
	"avyaas/utils"
)

/*
ResendUserOTP generates a new OTP and re-sends it to the user's phone number for verification.

Parameters:
  - phone: The user's phone number string.

Returns:
  - err: An error occured during the process, if any.
*/
func (repo *Repository) ResendUserOTP(identity string) error {
	// Generate a random OTP
	otp, err := utils.GenerateOTP()
	if err != nil {
		return err
	}

	// Create and save the new OTP and user's phone number to the database
	err = repo.SaveUserOTP(identity, otp)
	if err != nil {
		return err
	}

	if isValidEmail := utils.IsValidEmail(identity); isValidEmail {
		// Send the OTP to the user in provided email
		if err = utils.SendOTPEmail(identity, otp); err != nil {
			return err
		}
	} else if ok := utils.ContainsOnlyNumber(identity); ok {
		// Send the OTP to the user in provided phone number
		if err = utils.SendOTPSMS(identity, otp); err != nil {
			return err
		}
	}

	return nil
}
