package handler

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.FAQUsecase
}

func New(publicApp fiber.Router, privateApp fiber.Router, faqUsecase interfaces.FAQUsecase) {
	handler := &handler{
		usecase: faqUsecase,
	}
	publicFaqHandler := publicApp.Group("/faq/")
	privateFaqHandler := privateApp.Group("/faq/")

	privateFaqHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateFaq()))
	publicFaqHandler.Get("list/", handler.ListFaq())
	privateFaqHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateFaq()))
	privateFaqHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteFaq()))
}
