package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.FeedbackUsecase
}

func New(app fiber.Router, feedbackUsecase interfaces.FeedbackUsecase) {
	handler := &handler{
		usecase: feedbackUsecase,
	}

	feedbackHandler := app.Group("/feedback/")
	feedbackHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateFeedback()))
	feedbackHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListFeedback()))
	feedbackHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateFeedback()))
	feedbackHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteFeedback()))

}
