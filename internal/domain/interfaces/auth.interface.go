package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

/*
Usecase represents the authentication usecase interface, defining methods for handling various
authentication-related operations.
*/
type AuthUsecase interface {
	LoginUsecase(c *fiber.Ctx, user *presenter.LoginRequest) (presenter.LoginResponse, map[string]string)
	LogoutUsecase(c *fiber.Ctx) map[string]string
	GenerateAccessFromRefreshUsecase(user *presenter.AccessTokenRequest) (presenter.NewAccessTokenResponse, map[string]string)

	RegisterStudent(request presenter.StudentRegisterRequest) map[string]string
	VerifyUserOTP(otpRequest models.UserOtp) (bool, error)
	ResendUserOTP(identity string) error
	ForgotPassword(request presenter.ForgotPasswordRequest) error
	ResetPassword(request presenter.ResetPasswordRequest) map[string]string
}

/*
Repository represents the authentication repository interface, defining methods for handling various
authentication-related operations.
*/
type AuthRepository interface {
	GetRoleByID(roleID int) (*presenter.UserRoleResponse, error)

	RegisterStudent(user models.User, referral string, cID uint) error
	VerifyUserOTP(otpRequest models.UserOtp) (bool, error)
	ResendUserOTP(identity string) error
	CheckUserOTP(identity string) (*models.UserOtp, error)
	ForgotPassword(request presenter.ForgotPasswordRequest) error
	ResetPassword(request presenter.ResetPasswordRequest) error

	GetRegistrationNumber(transaction *gorm.DB, cID uint) (string, error)
}
