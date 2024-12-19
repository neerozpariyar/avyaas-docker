package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (u *usecase) ListReferral(request presenter.ReferralListRequest) ([]presenter.ReferralResponse, int, error) {
	referrals, totalPage, err := u.repo.ListReferral(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allReferrals []presenter.ReferralResponse

	for i := range referrals {
		ref, _ := u.repo.CheckPendingReferralInTransaction(referrals[i].ID)

		referral := presenter.ReferralResponse{
			ID:           referrals[i].ID,
			Title:        referrals[i].Title,
			Code:         referrals[i].Code,
			Type:         referrals[i].Type,
			DiscountType: referrals[i].DiscountType,
			Discount:     referrals[i].Discount,
			HasLimit:     referrals[i].HasLimit,
			HasUsed:      referrals[i].HasUsed,
			Limit:        ref.Limit,
		}

		if referrals[i].ExpiryDate != nil {
			referral.ExpiryDate = referrals[i].ExpiryDate.UTC().Format("2006-01-02T15:04:05Z")
		}

		if referrals[i].CourseID != 0 {
			course, err := u.courseRepo.GetCourseByID(referrals[i].CourseID)
			if err != nil {
				return nil, int(totalPage), err
			}

			courseData := make(map[string]interface{})
			courseData["id"] = course.ID
			courseData["courseID"] = course.CourseID

			referral.Course = courseData
		}

		allReferrals = append(allReferrals, referral)
	}

	return allReferrals, int(totalPage), nil
}
