package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type PaymentUsecase interface {
	ListPayments(request *presenter.PaymentListRequest) ([]presenter.PaymentListResponse, int, error)
	InitiateKhaltiPayment(request presenter.InitateKhaltiPaymentRequest) (*presenter.InitiateKhaltiPaymentResponse, map[string]string)
	VerifyPayment(request presenter.VerifyPaymentRequest) (*presenter.VerifyPaymentStatus, map[string]string)
	AddManualPayment(request presenter.ManualPaymentRequest) error // AddManualPayment adds a manual payment

	BulkAcessPayment(request *presenter.BulkAccessPaymentRequest) error
}

type PaymentRepository interface {
	ListPayments(request *presenter.PaymentListRequest) ([]models.Payment, float64, error)
	CreatePayment(request *models.Payment) error
	UpdatePayment(paymentID uint, request *models.Payment) error
	GetUserPaymentByCoursePackage(userID, courseID, packageID uint) (payment *models.Payment, err error)
	GetPaymentByTransactionID(transactionID string) (*models.Payment, error)
}
