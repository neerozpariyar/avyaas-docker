package http

import (
	"avyaas/internal/domain/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.BookmarkUsecase
}

func New(app fiber.Router, bookmarkUsecase interfaces.BookmarkUsecase) {
	handler := &handler{
		usecase: bookmarkUsecase,
	}

	bookmarkHandler := app.Group("/bookmark/")
	bookmarkHandler.Post("create/", handler.CreateBookmark())
	bookmarkHandler.Get("list/", handler.ListBookmark())
	// bookmarkHandler.Patch("update/:id/", handler.UpdateBookmark())
	bookmarkHandler.Get("details/:id/", handler.GetBookmarkDetails())
	bookmarkHandler.Delete("delete/:id/", handler.DeleteBookmark())

}
