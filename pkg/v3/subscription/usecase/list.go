package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) ListSubscriptions(request presenter.ListSubscriptionRequest) ([]presenter.SubscriptionListResponse, int, error) {
	subscriptions, totalPage, err := uCase.repo.ListSubscriptions(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allSubscriptions []presenter.SubscriptionListResponse

	for _, subscription := range subscriptions {
		eachSub := presenter.SubscriptionListResponse{
			ID:            subscription.ID,
			PaymentID:     subscription.PaymentID,
			PaymentMethod: subscription.PaymentMethod,
			TransactionID: subscription.TransactionID,
			ReferralCode:  subscription.ReferralCode,
		}

		course, err := uCase.courseRepo.GetCourseByID(subscription.CourseID)
		if err != nil {
			return nil, 0, err
		}

		courseData := make(map[string]interface{})
		courseData["id"] = course.ID
		courseData["title"] = course.Title
		eachSub.Course = courseData

		if subscription.ExpiryDate != nil {
			eachSub.ExpiryDate = subscription.ExpiryDate.UTC().Format("2006-01-02T15:04:05Z")
		}

		allSubscriptions = append(allSubscriptions, eachSub)
	}

	return allSubscriptions, int(totalPage), nil
}
