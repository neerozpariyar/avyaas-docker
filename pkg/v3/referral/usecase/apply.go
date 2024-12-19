package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
	"time"
)

func (uCase *usecase) ApplyReferral(request presenter.ApplyReferralRequest) (*presenter.ApplyReferralResponse, map[string]string) {
	var err error
	errMap := make(map[string]string)

	referral, err := uCase.repo.GetReferralByCode(request.Code)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	pkg, err := uCase.packageRepo.GetPackageByID(request.PackageID)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	if referral.CourseID != 0 {
		course, err := uCase.courseRepo.GetCourseByID(pkg.CourseID)
		if err != nil {
			errMap["error"] = err.Error()
			return nil, errMap
		}

		if referral.CourseID != pkg.CourseID {
			errMap["error"] = fmt.Errorf("referral code: '%s' is only valid for course: '%s'", referral.Code, course.CourseID).Error()
			return nil, errMap
		}
	}

	if referral.UserID != 0 {
		if referral.UserID != request.UserID {
			errMap["error"] = "invalid referral code"
		}
	}

	// Check if the user has already applied for this promotion or referral code
	if userReferral, _ := uCase.repo.CheckUserReferral(request.UserID, referral.ID); userReferral != nil {
		errMap["error"] = "referral code already applied"
		return nil, errMap
	}

	// Fetch the referral instance again as it might have updated in CheckUserReferral function
	referral, err = uCase.repo.GetReferralByCode(request.Code)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	if referral.ExpiryDate.Unix() < time.Now().Unix() {
		errMap["error"] = "expired referral code"
		return nil, errMap
	}

	if *referral.HasLimit && referral.Limit == 0 {
		errMap["error"] = "referral code has reached its limit"
		return nil, errMap
	}

	var discountedPrice float32
	if referral.DiscountType == "PERCENTAGE" {
		discountedPrice = float32(pkg.Price) - (float32(pkg.Price) * referral.Discount / 100)
	} else if referral.DiscountType == "FLAT" {
		discountedPrice = float32(pkg.Price) - referral.Discount
	}

	response := &presenter.ApplyReferralResponse{
		Price:           float32(pkg.Price),
		DiscountedPrice: discountedPrice,
	}

	if err = uCase.repo.CreateReferralInTransaction(request.UserID, referral.ID, request.PackageID); err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	return response, errMap
}
