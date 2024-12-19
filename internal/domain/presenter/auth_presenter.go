package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Presenter struct for storing custom claims data generated during token generation
type JwtCustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

/*
Presenter sturct for storing login request data.

Note: Identity field can be either email, username or phone of the user
*/
type LoginRequest struct {
	Identity string `json:"identity" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Presenter struct for storing refresh token for re-generating access token request
type AccessTokenRequest struct {
	Refresh string `json:"refresh" validate:"required"`
}

// ForgotPasswordRequest represents a payload struct for requesting password updates
type ForgotPasswordRequest struct {
	UserID uint   `json:"-"`
	Phone  string `json:"phone" validate:"required"`
}

type StudentRegisterRequest struct {
	Identity        string `json:"identity" validate:"required"`
	CollegeName     string `json:"collegeName"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
	Referral        string `json:"-"`
	CourseID        uint   `json:"courseID"`
}

type ResetPasswordRequest struct {
	Identity        string `json:"identity" validate:"required"`
	Otp             string `json:"otp" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}

// Presenter struct for storing new access token re-generate success response
type NewAccessTokenResponse struct {
	Access string `json:"access"`
}

// Presenter struct for storing success login response
type LoginResponse struct {
	Refresh string            `json:"refresh"`
	Access  string            `json:"access"`
	User    UserLoginResponse `json:"user"`
}

// Presenter struct for storing User login response data
type UserLoginResponse struct {
	ID         uint             `json:"id"`
	FirstName  string           `json:"firstName"`
	MiddleName string           `json:"middleName"`
	LastName   string           `json:"lastName"`
	Username   string           `json:"username"`
	Email      string           `json:"email"`
	Gender     string           `json:"gender"`
	Phone      string           `json:"phone"`
	Role       UserRoleResponse `json:"role"`
	Verified   bool             `json:"verified"`
	Image      string           `json:"image"`
}

// Presenter struct for storing User role data
type UserRoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

/*
LoginSuccessResponse creates a success response for login operations.

Parameters:
  - data: The LoginResponse containing data related to the successful login.

Returns:
  - map: A pointer to a fiber.Map representing the success response.
*/
func LoginSuccessResponse(data LoginResponse) *fiber.Map {
	return &fiber.Map{
		"success": true,
		"data":    data,
	}
}

/*
LogoutSuccessResponse creates a success response for logout operations.

Returns:
  - map: A pointer to a fiber.Map representing the success response.
*/
func LogoutSuccessResponse() *fiber.Map {
	return &fiber.Map{
		"success": true,
	}
}

/*
AuthErrorResponse creates an error response for authentication-related operations.

Parameters:
  - errMap: A map containing error messages related to authentication operations.

Returns:
  - map: A pointer to a fiber.Map representing the error response.
*/
func AuthErrorResponse(errMap map[string]string) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"errors":  errMap,
	}
}

/*
NewAccessTokenSuccessResponse creates a success response map for the "GenerateAccessTokenFromRefreshToken"
endpoint.

Parameters:
  - data: A NewAccessTokenResponse presenter struct containing the newly generated access token.

Returns:
  - map: A Fiber Map representing the success response with the provided data.
*/
func NewAccessTokenSuccessResponse(data NewAccessTokenResponse) *fiber.Map {
	return &fiber.Map{
		"success": true,
		"data":    data,
	}
}
