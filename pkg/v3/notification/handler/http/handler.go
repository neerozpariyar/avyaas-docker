package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.NotificationUsecase
}

func New(app fiber.Router, notificationUsecase interfaces.NotificationUsecase) {
	handler := &handler{
		usecase: notificationUsecase,
	}

	notificationHandler := app.Group("/notification/")
	notificationHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateNotification()))
	notificationHandler.Post("publish/:id/", middleware.RolesAndPermissionMiddleware(handler.PublishNotification()))

	notificationHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListNotification()))
	notificationHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateNotification()))
	notificationHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteNotification()))
	notificationHandler.Post("add-fcm-token/", middleware.RolesAndPermissionMiddleware(handler.AddFCMToken()))

}
