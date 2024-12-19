package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.SubjectUsecase
}

func New(app fiber.Router, subjectUsecase interfaces.SubjectUsecase) {
	handler := &handler{
		usecase: subjectUsecase,
	}

	subjectHandler := app.Group("/subject/")
	subjectHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateSubject()))
	subjectHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListSubject()))
	subjectHandler.Get("details/:id/", middleware.RolesAndPermissionMiddleware(handler.GetSubjectDetails()))
	subjectHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateSubject()))
	subjectHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteSubject()))
	subjectHandler.Patch("assign-units/", middleware.RolesAndPermissionMiddleware(handler.AssignUnitsToSubject()))
	subjectHandler.Get("get-heirarchy/:id/", middleware.RolesAndPermissionMiddleware(handler.GetSubjectHeirarchy()))
}
