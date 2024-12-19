package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"

	"fmt"
)

func (uc *usecase) AddManualPayment(request presenter.ManualPaymentRequest) error {
	pkg, err := uc.packageRepo.GetPackageByID(request.PackageID)
	if err != nil {
		return err
	}

	existingPayment, _ := uc.repo.GetPaymentByTransactionID(request.InvoiceNumber)
	if existingPayment != nil {
		return fmt.Errorf("payment with invoice number: '%s' already exists", request.InvoiceNumber)
	}

	payment := models.Payment{
		UserID:             request.UserID,
		PackageID:          request.PackageID,
		CourseID:           pkg.CourseID,
		Amount:             pkg.Price,
		Status:             "success",
		SubscriptionPeriod: pkg.Period,
		TransactionID:      request.InvoiceNumber,
	}

	subscription := presenter.SubscribePackageRequest{
		UserID:        request.UserID,
		PackageID:     request.PackageID,
		TransactionID: request.InvoiceNumber,
	}

	err = uc.repo.CreatePayment(&payment)
	if err != nil {
		return err
	}

	courseID, err := uc.packageRepo.GetCourseIDByPackageID(request.PackageID)
	if err != nil {
		return err
	}

	err = uc.courseRepo.EnrollInCourse(request.UserID, courseID) //if student is not enrolled in the course, enroll them
	if err != nil {
		if err.Error() == "already enrolled" { //if student is already enrolled in the course, update the subscription status
			err = uc.packageRepo.SubscribePackage(subscription)
			if err != nil {
				return err
			}
		}
		return err
	}

	err = uc.packageRepo.SubscribePackage(subscription)
	if err != nil {
		return err
	}

	return err
}
