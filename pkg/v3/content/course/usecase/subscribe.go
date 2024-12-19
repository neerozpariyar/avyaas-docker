package usecase

// import (
// 	"avyaas/internal/domain/presenter"
// )

// func (uCase *usecase) SubscribeCourse(request presenter.SubscribeCourseRequest) map[string]string {
// 	var err error

// 	errMap := make(map[string]string)
// 	_, err = uCase.repo.GetCourseByID(request.CourseID)
// 	if err != nil {
// 		errMap["error"] = err.Error()
// 		return errMap
// 	}

// 	packageData, err := uCase.packageRepo.GetPackageByID(request.PackageID)
// 	if err != nil {
// 		errMap["error"] = err.Error()
// 		return errMap
// 	}

// 	request.Period = packageData.Period

// 	_, err = uCase.packageRepo.CheckCoursePackage(request.CourseID, request.PackageID)
// 	if err != nil {
// 		errMap["error"] = err.Error()
// 		return errMap
// 	}

// 	payment, err := uCase.paymentRepo.GetUserPaymentByCoursePackage(request.UserID, request.CourseID, request.PackageID)
// 	if err != nil {
// 		errMap["error"] = err.Error()
// 	}

// 	if payment != nil && payment.Status != "COMPLETE" {
// 		errMap["error"] = "your payment is not complete"
// 		return errMap
// 	}

// 	request.PaymentID = payment.ID

// 	err = uCase.repo.SubscribeCourse(request)
// 	if err != nil {
// 		errMap["error"] = err.Error()
// 		return errMap
// 	}

// 	return errMap
// }
