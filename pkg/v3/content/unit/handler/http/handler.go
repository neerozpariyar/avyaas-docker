package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.UnitUsecase
}

func New(app fiber.Router, unitUsecase interfaces.UnitUsecase) {
	handler := &handler{
		usecase: unitUsecase,
	}

	unitHandler := app.Group("/unit/")
	unitHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateUnit()))
	unitHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListUnit()))
	unitHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateUnit()))
	unitHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteUnit()))

	unitHandler.Patch("update-position/", middleware.RolesAndPermissionMiddleware(handler.UpdateUnitPosition()))
	unitHandler.Patch("assign-chapters/", middleware.RolesAndPermissionMiddleware(handler.AssignChaptersToUnit()))
}
