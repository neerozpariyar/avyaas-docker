package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) SubscribePackage(request presenter.SubscribePackageRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	pkg, err := uCase.repo.GetPackageByID(request.PackageID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	_, err = uCase.repo.CheckCoursePackage(pkg.CourseID, request.PackageID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	payment, err := uCase.paymentRepo.GetUserPaymentByCoursePackage(request.UserID, pkg.CourseID, request.PackageID)
	if err != nil {
		errMap["error"] = err.Error()
	}

	if payment != nil && payment.Status != "COMPLETE" {
		errMap["error"] = "your payment is not complete"
		return errMap
	}
	if payment != nil {
		if payment.Status != "COMPLETE" {
			errMap["error"] = "your payment is not complete"
			return errMap
		}
		request.PaymentID = payment.ID
	}

	err = uCase.repo.SubscribePackage(request)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
