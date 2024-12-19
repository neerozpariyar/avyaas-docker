package models

import "time"

type Subscription struct {
	Timestamp

	UserID        uint
	CourseID      uint
	PackageID     uint
	PaymentID     uint
	PaymentMethod string
	TransactionID string
	ReferralCode  string
	ExpiryDate    *time.Time
}
