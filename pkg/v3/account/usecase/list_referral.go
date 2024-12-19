package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

func (uCase *usecase) ListTeacherReferrals(id uint) ([]presenter.TeacherReferralList, error) {

	var response []presenter.TeacherReferralList

	students, err := uCase.repo.GetStudentsByTeacherReferral(id)
	if err != nil {
		return nil, fmt.Errorf("teacher referral:'%d' doesn't exists", id)
	}

	for _, student := range students {
		user, err := uCase.repo.GetUserByID(student.ID)
		if err != nil {
			return nil, err
		}
		eachStudent := struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		}{
			ID:   student.ID,
			Name: user.FirstName,
		}

		subscriptions, _ := uCase.repo.GetSubscriptionByUserID(student.ID)
		if err == nil {
			continue
		}

		for _, subscription := range subscriptions {
			packageData, err := uCase.packageRepo.GetPackageByID(subscription.PackageID)
			if err != nil {
				return nil, err
			}
			eachSubscription := struct {
				Name          string  `json:"name"`
				Price         float64 `json:"price"`
				PaymentMethod string  `json:"paymentMethod"`
			}{
				Name:          packageData.Title,
				Price:         float64(packageData.Price),
				PaymentMethod: subscription.PaymentMethod,
			}

			eachResponse := presenter.TeacherReferralList{
				Student:      eachStudent,
				Subscription: eachSubscription,
			}

			response = append(response, eachResponse)
		}

	}
	return response, nil
}
