package handler

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.NoticeUsecase
}

func New(app fiber.Router, NoticeUsecase interfaces.NoticeUsecase) {
	handler := &handler{
		usecase: NoticeUsecase,
	}

	noticeHandler := app.Group("/notice/")
	noticeHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateNotice()))
	noticeHandler.Get("list/", handler.ListNotice())
	noticeHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteNotice()))
	noticeHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateNotice()))
}
