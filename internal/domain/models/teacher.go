package models

/*
Teacher represents a database model for storing teacher-related information, extending the Timestamp
model.
*/
type Teacher struct {
	Timestamp

	ReferralCode  string `json:"referralCode"`
	ReferralCount int64  `json:"referralCount" gorm:"default:0"`
}
