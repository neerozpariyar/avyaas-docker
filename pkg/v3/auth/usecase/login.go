package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"github.com/gofiber/fiber/v2"
)

/*
LoginUsecase performs the logic for user login, including validating credentials, creating access
and refresh tokens, and constructing the appropriate login response.

Parameters:
  - c: Fiber context representing the incoming HTTP request.
  - user: Pointer to an presenter.LoginRequest presenter containing user login information.

Returns:
  - response: An presenter.LoginResponse presener containing access and refresh tokens, along with user details.
  - errMap: A map containing error details, if any occurred during the login process.
*/
func (uCase *usecase) LoginUsecase(c *fiber.Ctx, user *presenter.LoginRequest) (presenter.LoginResponse, map[string]string) {
	errMap := make(map[string]string)
	var response presenter.LoginResponse

	// Retrieve the user by email
	result, err := uCase.accountRepo.GetUserByEmail(user.Identity)
	if err != nil {
		errMap["error"] = "invalid credentials"
	} else {
		errMap = make(map[string]string)
	}

	// If no user is found by email, attempt to retrieve a user by username
	if result == nil {
		result, err = uCase.accountRepo.GetUserByUsername(user.Identity)
		if err != nil {
			errMap["error"] = "invalid credentials"
		} else {
			errMap = make(map[string]string)
		}
	}

	// If no user is found by email, attempt to retrieve a user by phone
	if result == nil {
		result, err = uCase.accountRepo.GetUserByPhone(user.Identity)
		if err != nil {
			errMap["error"] = "invalid credentials"
		} else {
			errMap = make(map[string]string)
		}
	}

	if len(errMap) != 0 {
		return response, errMap
	}

	errMap = utils.ValidateAccess(c.Get("Origin"), result.RoleID)
	if len(errMap) != 0 {
		return response, errMap
	}

	// Check if the provided password matches the stored hashed password
	if !utils.CheckPasswordHash(user.Password, result.Password) {
		errMap := make(map[string]string)
		errMap["error"] = "invalid credentials"
		return response, errMap
	}

	// Check if the user account is verified
	if !result.Verified {
		errMap := make(map[string]string)

		// Resend the OTP for verification
		err = uCase.repository.ResendUserOTP(user.Identity)
		if err != nil {
			errMap["otp"] = err.Error()
		}

		errMap["error"] = "unverified account: please verify your account"
		return response, errMap
	}

	// Generate an access token for the authenticated user
	accessToken, err := uCase.CreateAccessToken(result.ID, result.Username)
	if err != nil {
		errMap["error"] = err.Error()
		return response, errMap
	}

	// Generate a refresh token for the authenticated user
	refreshToken, err := uCase.CreateRefreshToken(result.ID, result.Username)
	if err != nil {
		errMap["error"] = err.Error()
		return response, errMap
	}

	// Retrieve the user's role information
	role, err := uCase.repository.GetRoleByID(result.RoleID)
	if err != nil {
		errMap["error"] = err.Error()
		return response, errMap
	}

	// Construct the login response with access and refresh tokens, along with user details
	response = presenter.LoginResponse{
		Refresh: refreshToken,
		Access:  accessToken,
		User: presenter.UserLoginResponse{
			ID:         result.ID,
			FirstName:  result.FirstName,
			MiddleName: result.MiddleName,
			LastName:   result.LastName,
			Username:   result.Username,
			Email:      result.Email,
			Gender:     result.Gender,
			Phone:      result.Phone,
			Role:       *role,
			Verified:   result.Verified,
		},
	}

	return response, errMap
}
