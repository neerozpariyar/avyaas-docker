package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

/*
handler represents the HTTP handler for the test module, providing methods to handle various HTTP
requests related to test using the specified usecase.
*/
type handler struct {
	usecase interfaces.TestUsecase
}

/*
New initializes and configures the test module within the Fiber app. It creates a test service
handler with the provided usecase and sets up routes for various operations related to the test
under the specified base path.
*/
func New(app fiber.Router, usecase interfaces.TestUsecase) {
	// Create a test service handler with the provided usecase
	handler := &handler{
		usecase: usecase,
	}

	testTypeHandler := app.Group("/test-type/")
	testTypeHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateTestType()))
	testTypeHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListTestType()))
	testTypeHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateTestType()))
	testTypeHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteTestType()))

	testHandler := app.Group("/test/")
	testHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateTest()))
	testHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListTest()))
	testHandler.Get("details/:id/", middleware.RolesAndPermissionMiddleware(handler.GetTestDetails()))
	testHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateTest()))
	testHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteTest()))
	testHandler.Post("assign-question-set/", middleware.RolesAndPermissionMiddleware(handler.AssignQuestionSetToTest()))
	testHandler.Patch("update-status/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateTestStatus()))

	testHandler.Post("submit/:id/", middleware.RolesAndPermissionMiddleware(handler.SubmitTest()))
	testHandler.Get("leaderboard/:id/", middleware.RolesAndPermissionMiddleware(handler.GetTestLeaderboard()))
	testHandler.Get("result/:id/", middleware.RolesAndPermissionMiddleware(handler.GetTestResult()))
	testHandler.Get("history/", middleware.RolesAndPermissionMiddleware(handler.GetTestHistory()))
}

/* Note: The files for test services are named as quiz in suffix. For example: create_quiz as create_test
will refer to a test case by default. */
