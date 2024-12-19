package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.ContentUsecase
}

func New(app fiber.Router, ContentUsecase interfaces.ContentUsecase) {
	handler := &handler{
		usecase: ContentUsecase,
	}

	contentHandler := app.Group("/content/")
	contentHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateContent()))
	contentHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListContent()))
	contentHandler.Get("details/:id/", middleware.RolesAndPermissionMiddleware(handler.GetContentDetails()))
	contentHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateContent()))
	contentHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteContent()))
	contentHandler.Patch("assign/", middleware.RolesAndPermissionMiddleware(handler.AssignContentsToRelation()))
	contentHandler.Patch("update-position/", middleware.RolesAndPermissionMiddleware(handler.UpdateContentPosition()))
	//progress

	contentHandler.Patch("mark-completed/:id/", middleware.RolesAndPermissionMiddleware(handler.MarkAsCompleted()))
	contentHandler.Patch("update-progress/", middleware.RolesAndPermissionMiddleware(handler.EvaluateContentProgress())) //for content

	contentHandler.Post("find-length/", middleware.RolesAndPermissionMiddleware(handler.FindVideoLength())) //for content
}
