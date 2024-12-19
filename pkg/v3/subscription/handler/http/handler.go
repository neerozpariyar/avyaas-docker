package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.SubscriptionUsecase
}

func New(app fiber.Router, subscriptionUsecase interfaces.SubscriptionUsecase) {
	handler := &handler{
		usecase: subscriptionUsecase,
	}

	subscriptionHandler := app.Group("/subscription/")
	subscriptionHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListSubscriptions()))
}
