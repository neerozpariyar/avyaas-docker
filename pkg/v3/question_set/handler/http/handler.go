package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

/*
handler represents the HTTP handler for the question set module, providing methods to handle various
HTTP requests related to question set using the specified usecase.
*/
type handler struct {
	usecase interfaces.QuestionSetUsecase
}

/*
New initializes and configures the question set module within the Fiber app. It creates a question
set service handler with the provided usecase and sets up routes for various operations related to
the question set under the specified base path.
*/
func New(app fiber.Router, usecase interfaces.QuestionSetUsecase) {
	// Create a question set service handler with the provided usecase
	handler := &handler{
		usecase: usecase,
	}

	qsHandler := app.Group("/question-set/")
	qsHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateQuestionSet()))
	qsHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListQuestionSet()))
	qsHandler.Get("details/:id/", middleware.RolesAndPermissionMiddleware(handler.GetQuestionSetDetails()))
	qsHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateQuestionSet()))
	qsHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteQuestionSet()))
	qsHandler.Post("assign-questions/", middleware.RolesAndPermissionMiddleware(handler.AssignQuestionsToQuestionSet()))
}
