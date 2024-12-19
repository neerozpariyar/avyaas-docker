package usecase

import "avyaas/internal/domain/models"

/*
VerifyUserOTP is a usecase method that verifies a user's OTP (One-Time Password).

Parameters:
  - otpRequest: A models.UserOtp instance as a parameter.

Returns:
  - verified: A boolean indicating whether the OTP is valid
  - err: An error, if any, encountered during the verification process.
*/
func (uCase *usecase) VerifyUserOTP(otpRequest models.UserOtp) (bool, error) {
	verified, err := uCase.repository.VerifyUserOTP(otpRequest)

	return verified, err
}
