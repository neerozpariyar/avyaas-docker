package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

/*
handler represents the HTTP handler for the course group module, providing methods to handle various
HTTP requests related to course group using the specified usecase.
*/
type handler struct {
	usecase interfaces.CourseGroupUsecase
}

/*
New initializes and configures the course group module within the Fiber app. It creates a course group
service handler with the provided usecase and sets up routes for various operations related to the
course groups under the specified base path.
*/
func New(app fiber.Router, cgUsecase interfaces.CourseGroupUsecase) {
	// Create a course group service handler with the provided usecase
	handler := &handler{
		usecase: cgUsecase,
	}

	cgHandler := app.Group("/course-group/")
	cgHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateCourseGroup()))
	cgHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListCourseGroup()))
	cgHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateCourseGroup()))
	cgHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteCourseGroup()))
	cgHandler.Patch("assign-course/:id/", middleware.RolesAndPermissionMiddleware(handler.AssignCoursesToCourseGroup()))
}
