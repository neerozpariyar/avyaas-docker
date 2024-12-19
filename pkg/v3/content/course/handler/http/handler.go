package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.CourseUsecase
}

/*
New initializes and configures the course  module within the Fiber app. It creates a course
service handler with the provided usecase and sets up routes for various operations related to the
course under the specified base path.
*/
func New(app fiber.Router, courseUsecase interfaces.CourseUsecase) {
	// Create a course  service handler with the provided usecase
	handler := &handler{
		usecase: courseUsecase,
	}

	courseHandler := app.Group("/course/")
	courseHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateCourse()))
	courseHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListCourse()))
	courseHandler.Get("details/:id/", middleware.RolesAndPermissionMiddleware(handler.GetCourseDetails()))

	courseHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateCourse()))
	courseHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteCourse()))
	courseHandler.Post("enroll/:id/", middleware.RolesAndPermissionMiddleware(handler.EnrollUser()))
	// courseHandler.Post("subscribe/:id/", middleware.RolesAndPermissionMiddleware(handler.SubscribeCourse()))
	courseHandler.Get("enrolled/", middleware.RolesAndPermissionMiddleware(handler.ListEnrolledCourse()))
	courseHandler.Patch("update-availability/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateAvailability()))
	courseHandler.Patch("assign-subject/:id/", middleware.RolesAndPermissionMiddleware(handler.AssignSubjectsToCourse()))

}
