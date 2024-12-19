package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.LiveGroupUsecase
}

func New(app fiber.Router, LiveGroupUsecase interfaces.LiveGroupUsecase) {
	handler := &handler{
		usecase: LiveGroupUsecase,
	}

	liveGroupHandler := app.Group("/live-group/")
	liveGroupHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateLiveGroup()))
	liveGroupHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListLiveGroup()))
	// liveGroupHandler.Get("details/:id/", middleware.RolesAndPermissionMiddleware(handler.GetLiveGroupDetails()))
	liveGroupHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateLiveGroup()))
	liveGroupHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteLiveGroup()))

}
