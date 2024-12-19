package models

/*
Student represents a database model for storing student-related information, extending the Timestamp
model.
*/
type Student struct {
	Timestamp
	ReferralCode string `json:"referralCode"`
	//ReferralCount      int64  `json:"referralCount"`
	ReferredBy         uint   `json:"referredBy"`
	RegistrationNumber string `json:"registrationNumber" gorm:"unique"`
}
