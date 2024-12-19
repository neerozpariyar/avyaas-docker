package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

/*
handler represents the HTTP handler for the service module, providing methods to handle various HTTP
requests related to service using the specified usecase.
*/
type handler struct {
	usecase interfaces.ServiceUsecase
}

/*
New initializes and configures the service module within the Fiber app. It creates a service handler
with the provided usecase and sets up routes for various operations related to the services under the
specified base path.
*/
func New(app fiber.Router, usecase interfaces.ServiceUsecase) {
	// Create an service handler with the provided usecase
	handler := &handler{
		usecase: usecase,
	}

	serviceHandler := app.Group("/service/")
	serviceHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateService()))
	serviceHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListService()))
	serviceHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateService()))
	serviceHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteService()))
}
