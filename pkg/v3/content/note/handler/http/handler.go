package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.NoteUsecase
}

func New(app fiber.Router, NoteUsecase interfaces.NoteUsecase) {
	handler := &handler{
		usecase: NoteUsecase,
	}

	chapterHandler := app.Group("/note/")
	chapterHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateNote()))
	chapterHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListNote()))
	chapterHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateNote()))
	chapterHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteNote()))
}
