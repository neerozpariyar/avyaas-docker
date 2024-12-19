package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

/*
handler represents the HTTP handler for the question module, providing methods to handle various
HTTP requests related to question using the specified usecase.
*/
type handler struct {
	usecase interfaces.TypeQuestionUsecase
}

/*
New initializes and configures the question module within the Fiber app. It creates a question
service handler with the provided usecase and sets up routes for various operations related to the
question under the specified base path.
*/
func New(app fiber.Router, usecase interfaces.TypeQuestionUsecase) {
	// Create a question set service handler with the provided usecase
	handler := &handler{
		usecase: usecase,
	}

	qHandler := app.Group("/tquestion/")
	qHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreateAllQuestion()))
	// qHandler.Post("create-case/", middleware.RolesAndPermissionMiddleware(handler.CreateCaseBasedQuestion()))
	qHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListTypeQuestion()))
	qHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdateAllQuestion()))
	qHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeleteTypeQuestion()))
}
