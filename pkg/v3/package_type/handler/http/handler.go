package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.PackageTypeUsecase
}

func New(app fiber.Router, usecase interfaces.PackageTypeUsecase) {
	handler := &handler{
		usecase: usecase,
	}

	packageTypeHandler := app.Group("/package-type/")
	packageTypeHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreatePackageType()))
	packageTypeHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListPackageType()))
	packageTypeHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdatePackageType()))
	packageTypeHandler.Post("delete/:id", middleware.RolesAndPermissionMiddleware(handler.DeletePackageType()))
}
