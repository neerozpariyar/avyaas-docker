package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

/*
handler represents the HTTP handler for the package module, providing methods to handle various HTTP
requests related to package using the specified usecase.
*/
type handler struct {
	usecase            interfaces.PackageUsecase
	packageTypeUsecase interfaces.PackageTypeUsecase
}

/*
New initializes and configures the package module within the Fiber app. It creates a package service
handler with the provided usecase and sets up routes for various operations related to the packages
under the specified base path.
*/
func New(app fiber.Router, usecase interfaces.PackageUsecase, packageTypeUsecase interfaces.PackageTypeUsecase) {
	// Create an package service handler with the provided usecase
	handler := &handler{
		usecase:            usecase,
		packageTypeUsecase: packageTypeUsecase,
	}

	packageHandler := app.Group("/package/")
	packageHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreatePackage()))
	packageHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListPackage()))
	packageHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdatePackage()))
	packageHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeletePackage()))
	packageHandler.Post("subscribe/:id/", middleware.RolesAndPermissionMiddleware(handler.SubscribePackage()))
}
