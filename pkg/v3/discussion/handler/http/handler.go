package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.DiscussionUsecase
}

func New(app fiber.Router, DiscussionUsecase interfaces.DiscussionUsecase) {
	handler := &handler{
		usecase: DiscussionUsecase,
	}

	discussionHandler := app.Group("/discussion/")
	discussionHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateDiscussion()))
	discussionHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListDiscussion()))
	// discussionHandler.Get("details/:id/", middleware.RolesAndPermissionMiddleware(handler.GetDiscussionDetails()))
	discussionHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateDiscussion()))
	discussionHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteDiscussion()))
	discussionHandler.Patch("vote/:id/", middleware.RolesAndPermissionMiddleware(handler.LikeOrUnlikeDiscussion()))

	// discussionHandler.Patch("update-position/", middleware.RolesAndPermissionMiddleware(handler.UpdateDiscussionPosition()))
}
