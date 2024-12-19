package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"strings"
)

func (uCase *usecase) VerifyPayment(request presenter.VerifyPaymentRequest) (*presenter.VerifyPaymentStatus, map[string]string) {
	var err error

	errMap := make(map[string]string)
	_, err = uCase.courseRepo.GetCourseByID(request.CourseID)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	packageData, err := uCase.packageRepo.GetPackageByID(request.PackageID)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	_, err = uCase.packageRepo.CheckCoursePackage(request.CourseID, request.PackageID)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	var response *presenter.VerifyPaymentStatus

	merchant := strings.ToUpper(request.PaymentMethod)
	switch merchant {
	case "ESEWA":
		_, response, err = verifyEsewaPayment(request.Amount, request.TransactionUUID, request.TransactionID)
		if err != nil {
			errMap["error"] = err.Error()
			return nil, errMap
		}
	case "KHALTI":
		_, response, err = verifyKhaltiPayment(request.TransactionUUID)
		if err != nil {
			errMap["error"] = err.Error()
			return nil, errMap
		}

		if response == nil {
			errMap["error"] = "error validating payment"
			return response, errMap
		}
	default:
		errMap["error"] = "invalid payment method"
		return nil, errMap
	}

	if strings.ToUpper(response.Status) == "COMPLETE" || strings.ToUpper(response.Status) == "COMPLETED" {
		paymentData := &models.Payment{
			UserID:             request.UserID,
			CourseID:           request.CourseID,
			PackageID:          request.PackageID,
			MerchantType:       merchant,
			TransactionID:      request.TransactionID,
			Amount:             request.Amount,
			SubscriptionPeriod: packageData.Period,
			Status:             "COMPLETE",
			TransactionUUID:    request.TransactionUUID,
		}

		payment, err := uCase.repo.GetUserPaymentByCoursePackage(request.UserID, request.CourseID, request.PackageID)
		if payment != nil && err == nil {
			err = uCase.repo.UpdatePayment(payment.ID, paymentData)
			if err != nil {
				errMap["error"] = err.Error()
				return nil, errMap
			}
		} else {
			err = uCase.repo.CreatePayment(paymentData)
			if err != nil {
				errMap["error"] = err.Error()
				return nil, errMap
			}
		}

		subscribeRequest := presenter.SubscribePackageRequest{
			UserID:          request.UserID,
			PackageID:       request.PackageID,
			PaymentMethod:   request.PaymentMethod,
			Amount:          request.Amount,
			TransactionID:   request.TransactionID,
			TransactionUUID: request.TransactionUUID,
		}

		errMap = uCase.packageUsecase.SubscribePackage(subscribeRequest)
		if len(errMap) != 0 {
			return response, errMap
		}

		return response, errMap
	}

	errMap["error"] = "error validating payment"
	return nil, errMap

}
