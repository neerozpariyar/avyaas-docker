package models

import "time"

/*
UserOtp is a struct representing a user's OTP (One-Time Password) information. It includes details
such as the phone number, OTP code, and expiration time.
*/
type UserOtp struct {
	Timestamp

	Identity  string
	OTP       string
	ExpiresAt time.Time
}
