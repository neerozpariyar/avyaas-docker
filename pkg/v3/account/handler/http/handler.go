package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

/*
handler represents the HTTP handler for the account module, providing methods to handle various HTTP
requests related to account using the specified usecase.
*/
type handler struct {
	usecase interfaces.AccountUsecase
}

/*
New initializes and configures the account module within the Fiber app. It creates a account service
handler with the provided usecase and sets up routes for various operations related to the accounts
under the specified base path.
*/
func New(app fiber.Router, usecase interfaces.AccountUsecase) {
	// Create an account service handler with the provided usecase
	handler := &handler{
		usecase: usecase,
	}

	accountHandler := app.Group("/account/")
	accountHandler.Patch("change-password/", middleware.RolesAndPermissionMiddleware(handler.ChangePassword()))

	teacherHandler := app.Group("/teacher/")
	teacherHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateTeacher()))
	teacherHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListTeacher()))
	teacherHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateTeacher()))
	teacherHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteTeacher()))
	teacherHandler.Patch("assign-subjects/:userID/", middleware.RolesAndPermissionMiddleware(handler.AssignSubjectToTeacher()))
	teacherHandler.Get("referrals/:id/", middleware.RolesAndPermissionMiddleware(handler.ListTeacherReferrals()))

	studentHandler := app.Group("/student/")
	studentHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListStudent()))
	studentHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateStudent()))
}
