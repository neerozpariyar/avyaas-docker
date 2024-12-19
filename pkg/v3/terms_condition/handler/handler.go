package handler

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.TermsAndConditionUsecase
}

func New(publicApp fiber.Router, privateApp fiber.Router, TermsAndConditionUsecase interfaces.TermsAndConditionUsecase) {
	handler := &handler{
		usecase: TermsAndConditionUsecase,
	}

	publicTermsHandler := publicApp.Group("/terms-condition/")
	privateTermsHandler := privateApp.Group("/terms-condition/")

	privateTermsHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateTermsAndCondition()))
	publicTermsHandler.Get("list/", handler.ListTermsAndCondition())
	privateTermsHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateTermsAndCondition()))
	privateTermsHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteTermsAndCondition()))
}
