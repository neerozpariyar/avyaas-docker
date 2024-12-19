package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.LiveUsecase
}

func New(app fiber.Router, LiveUsecase interfaces.LiveUsecase) {
	handler := &handler{
		usecase: LiveUsecase,
	}

	liveHandler := app.Group("/live/")
	liveHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateLive()))
	liveHandler.Post("msdk/", middleware.RolesAndPermissionMiddleware(handler.MeetingSDKKey()))

	liveHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListLive()))
	// liveHandler.Get("details/:id/", middleware.RolesAndPermissionMiddleware(handler.GetLiveDetails()))
	liveHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateLive()))
	liveHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteLive()))
}
