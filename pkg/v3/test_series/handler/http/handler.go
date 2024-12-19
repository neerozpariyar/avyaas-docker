package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.TestSeriesUsecase
}

func New(app fiber.Router, testSeriesUsecase interfaces.TestSeriesUsecase) {
	handler := &handler{
		usecase: testSeriesUsecase,
	}

	tsHandler := app.Group("/test-series/")
	tsHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateTestSeries()))
	tsHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListTestSeries()))
	tsHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateTestSeries()))
	tsHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteTestSeries()))
}
