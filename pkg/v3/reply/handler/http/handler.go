package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.ReplyUsecase
}

func New(app fiber.Router, ReplyUsecase interfaces.ReplyUsecase) {
	handler := &handler{
		usecase: ReplyUsecase,
	}

	discussionHandler := app.Group("/reply/")
	discussionHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateReply()))
	discussionHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListReply()))
	// discussionHandler.Get("details/:id/", middleware.RolesAndPermissionMiddleware(handler.GetReplyDetails()))
	discussionHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateReply()))
	discussionHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteReply()))
	// discussionHandler.Patch("update-position/", middleware.RolesAndPermissionMiddleware(handler.UpdateReplyPosition()))
}
