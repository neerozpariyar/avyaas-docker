package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

/*
handler represents the HTTP handler for the package module, providing methods to handle various HTTP
requests related to package using the specified usecase.
*/
type handler struct {
	usecase interfaces.PaymentUsecase
}

/*
New initializes and configures the package module within the Fiber app. It creates a package service
handler with the provided usecase and sets up routes for various operations related to the packages
under the specified base path.
*/
func New(app fiber.Router, usecase interfaces.PaymentUsecase) {
	// Create an package service handler with the provided usecase
	handler := &handler{
		usecase: usecase,
	}

	paymentHandler := app.Group("/payment/")
	paymentHandler.Post("initiate-khalti/", middleware.RolesAndPermissionMiddleware(handler.InitiateKhaltiPayment()))
	paymentHandler.Post("verify/", middleware.RolesAndPermissionMiddleware(handler.VerifyPayment()))
	paymentHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListPayments()))
	paymentHandler.Post("add-manual-payment/", middleware.RolesAndPermissionMiddleware(handler.AddManualPayment()))
	paymentHandler.Post("bulk-access-payment/", middleware.RolesAndPermissionMiddleware(handler.BulkAccessPaymentRequest()))
}
