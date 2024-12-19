package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.FileClientUsecase
}

func New(app fiber.Router, FileClientUseCase interfaces.FileClientUsecase) {
	handler := &handler{
		usecase: FileClientUseCase,
	}

	fileClientHandler := app.Group("/file-client/")

	fileClientHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListObjects()))

	fileClientHandler.Delete("delete/", middleware.RolesAndPermissionMiddleware(handler.DeleteObjects()))
}
