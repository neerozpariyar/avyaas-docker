package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"encoding/json"
	"time"
)

func (uCase *usecase) UpdateReferral(data presenter.CreateUpdateReferralRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing content  with the provided content 's ID
	_, err = uCase.repo.GetReferralByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

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

	var refData models.Referral

	bData, err := json.Marshal(data)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	err = json.Unmarshal(bData, &refData)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Parse and set the string type end time to *time.Time if provided
	var et time.Time
	if data.ExpiryDate != "" {
		expiryDate := data.ExpiryDate
		if et, err = time.Parse(time.RFC3339, expiryDate); err != nil {
			errMap["expiryDate"] = "error parsing invalid UTC time"
			return errMap
		}

		refData.ExpiryDate = &et
	}

	// Delegate the update of content
	if err = uCase.repo.UpdateReferral(refData); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
