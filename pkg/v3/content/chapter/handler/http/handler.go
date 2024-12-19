package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.ChapterUsecase
}

func New(app fiber.Router, chapterUsecase interfaces.ChapterUsecase) {
	handler := &handler{
		usecase: chapterUsecase,
	}

	chapterHandler := app.Group("/chapter/")
	chapterHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateChapter()))
	chapterHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListChapter()))
	chapterHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateChapter()))
	chapterHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteChapter()))

	chapterHandler.Patch("update-position/", middleware.RolesAndPermissionMiddleware(handler.UpdateChapterPosition()))
}
