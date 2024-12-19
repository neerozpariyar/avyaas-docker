package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.ReferralUsecase
}

func New(app fiber.Router, ReferralUsecase interfaces.ReferralUsecase) {
	handler := &handler{
		usecase: ReferralUsecase,
	}

	discussionHandler := app.Group("/referral/")
	discussionHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateReferral()))
	discussionHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListReferral()))
	discussionHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateReferral()))
	discussionHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteReferral()))
	discussionHandler.Post("apply/", middleware.RolesAndPermissionMiddleware(handler.ApplyReferral()))
}
