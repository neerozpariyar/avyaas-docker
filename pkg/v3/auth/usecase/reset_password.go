package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (uCase *usecase) ResetPassword(request presenter.ResetPasswordRequest) map[string]string {
	errMap := make(map[string]string)
	var err error

	if isValidEmail := utils.IsValidEmail(request.Identity); isValidEmail {
		_, err := uCase.accountRepo.GetUserByEmail(request.Identity)
		if err != nil {
			errMap["user"] = err.Error()
			return errMap
		}
	} else if ok := utils.ContainsOnlyNumber(request.Identity); ok {
		_, err := uCase.accountRepo.GetUserByPhone(request.Identity)
		if err != nil {
			errMap["user"] = err.Error()
			return errMap
		}
	} else {
		errMap["identity"] = "invalid phone or email"
		return errMap
	}

	verifyData := models.UserOtp{
		Identity: request.Identity,
		OTP:      request.Otp,
	}

	if _, err := uCase.repository.VerifyUserOTP(verifyData); err != nil {
		errMap["otp"] = err.Error()
		return errMap
	}

	// Check the strength of the provided password
	err = utils.CheckPasswordStrength(request.NewPassword)
	if err != nil {
		errMap["password"] = err.Error()
		return errMap
	}

	// Set the request userID with the retrieved user's ID
	err = uCase.repository.ResetPassword(request)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
