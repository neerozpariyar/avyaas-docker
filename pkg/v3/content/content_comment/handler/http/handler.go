package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.ContentCommentUsecase
}

func New(app fiber.Router, ContentCommentUsecase interfaces.ContentCommentUsecase) {
	handler := &handler{
		usecase: ContentCommentUsecase,
	}

	replyHandler := app.Group("/content-comment/")
	replyHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateComment()))
	replyHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListComment()))
	// replyHandler.Get("details/:id/", middleware.RolesAndPermissionMiddleware(handler.GetCommentDetails()))
	replyHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateComment()))
	replyHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteComment()))
	// replyHandler.Patch("update-position/", middleware.RolesAndPermissionMiddleware(handler.UpdateCommentPosition()))
}
