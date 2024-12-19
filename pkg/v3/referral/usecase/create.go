package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"encoding/json"
	"time"
)

func (uCase *usecase) CreateReferral(data presenter.CreateUpdateReferralRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	if data.CourseID != 0 {
		if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
			errMap["courseID"] = err.Error()
			return errMap
		}
	}

	if data.UserID != 0 {
		if _, err := uCase.accountRepo.GetUserByID(data.UserID); err != nil {
			errMap["userID"] = err.Error()
			return errMap
		}
	}

	var referral models.Referral

	bData, err := json.Marshal(data)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	err = json.Unmarshal(bData, &referral)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}
	if data.DiscountType == "percentage" {
		if data.Discount > 100 {
			errMap["discountValue"] = "Discount value should be less than 100"
			return errMap
		}
	}
	// Parse and set the string type end time to *time.Time if provided
	var et time.Time
	if data.ExpiryDate != "" {
		expiryDate := data.ExpiryDate
		if et, err = time.Parse(time.RFC3339, expiryDate); err != nil {
			errMap["expiryDate"] = "error parsing invalid UTC time"
			return errMap
		}

		referral.ExpiryDate = &et
	}

	if err = uCase.repo.CreateReferral(referral); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
