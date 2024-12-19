package presenter

import "time"

/*
UserCreateUpdateRequest is a struct representing the request payload for creating or updating a user.
It includes fields such as ID, name, email, phone, role ID, and password.
*/
type UserCreateUpdateRequest struct {
	ID         uint   `json:"id"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	RoleID     int    `json:"-"`
	Password   string `json:"password" validate:"required"`
}

/*
VerifyOTPRequest is a struct representing the request payload for verifying an OTP (One-Time Password).
It includes fields such as Phone (phone number to verify) and OTP (the One-Time Password to be verified).
*/
type VerifyOTPRequest struct {
	Identity string `json:"identity" validate:"required"`
	OTP      string `json:"otp" validate:"required"`
}

// RequestOTPRequest represents a payload struct for resending OTP to the user's phone number.
type RequestOTPRequest struct {
	Identity string `json:"identity" validate:"required"`
}

// ChangePasswordRequest represents a payload struct for changing user's password
type ChangePasswordRequest struct {
	UserID          uint   `json:"-"`
	CurrentPassword string `json:"currentPassword" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required"`
	ConfirmPassword string `json:"confirmPassword"  validate:"required"`
}

// Presenter struct for storing User response data from database query
type UserResponse struct {
	CreatedAt  time.Time `json:"createdAt"`
	ID         uint      `json:"id"`
	FirstName  string    `json:"firstName"`
	MiddleName string    `json:"middleName"`
	LastName   string    `json:"lastName"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Gender     string    `json:"gender"`
	Phone      string    `json:"phone"`
	RoleID     int       `json:"roleID"`
	Verified   bool      `json:"verified"`
	Image      string    `json:"image"`
	Password   string    `json:"-"`
}

type UserReqPresenter struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"firstName"`
	MiddleName   string `json:"middleName"`
	LastName     string `json:"lastName"`
	Username     string `json:"username"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	RoleID       int    `json:"-"`
	Verified     bool   `gorm:"default:false" json:"-"`
	Image        string `json:"image"`
	Password     string `json:"password"`
	CollegeName  string `json:"collegeName"`
	OauthID      string `json:"-"`
	FacebookID   string `json:"-"`
	ReferralCode string `json:"referralCode"`
}
