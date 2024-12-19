package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.RecordingUsecase
}

func New(app fiber.Router, RecordingUsecase interfaces.RecordingUsecase) {
	handler := &handler{
		usecase: RecordingUsecase,
	}

	recordingHandler := app.Group("/recording/")
	recordingHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.UploadRecording()))
	recordingHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListRecording()))
	recordingHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateRecording()))
	recordingHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteRecording()))
}
