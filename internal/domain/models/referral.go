package models

import "time"

type Referral struct {
	Timestamp

	Title        string     `json:"title" validate:"required"`
	Type         string     `json:"type" validate:"required"` // Options: General, Student or Course
	CourseID     uint       `json:"courseID"`                 // should be required if referral type is: Course
	UserID       uint       `json:"userID"`                   // ID of student  who received the referral coupon
	Code         string     `json:"code" validate:"required"`
	ExpiryDate   *time.Time `json:"expiryDate" validate:"required"`
	DiscountType string     `json:"discountType" validate:"required"`
	Discount     float32    `json:"discount" validate:"required"`
	HasLimit     *bool      `json:"hasLimit" gorm:"default:false"`
	HasUsed      *bool      `json:"hasUsed" gorm:"default:false"`
	Limit        uint       `json:"limit" gorm:"default:0"`
}

// CheckAndUpdateHasUsed checks whether the referral limit has been reached and updates HasUsed accordingly.
// func (r *Referral) CheckAndUpdateHasUsed(db *gorm.DB) error {
// 	if r.HasLimit == nil || !*r.HasLimit {
// 		return nil // No limit, no need to update HasUsed
// 	}

// 	// Count the number of users who have used this referral code
// 	var usedCount int64
// 	if err := db.Model(&Referral{}).Where("code = ?", r.Code).Count(&usedCount).Error; err != nil {
// 		return err
// 	}

// 	if uint(usedCount) >= r.Limit {
// 		// update HasUsed to true
// 		if err := db.Save(r).Error; err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

type UserReferral struct {
	UserID     uint
	ReferralID uint
}

type ReferralInTransaction struct {
	Timestamp

	UserID     uint
	PackageID  uint
	ReferralID uint
	HoldTime   *time.Time
}
