package http

import (
	"avyaas/internal/domain/interfaces"
	"avyaas/pkg/v3/auth/handler/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.PollUsecase
}

func New(app fiber.Router, PollUsecase interfaces.PollUsecase) {
	handler := &handler{
		usecase: PollUsecase,
	}

	pollHandler := app.Group("/poll/")
	pollHandler.Post("create/", middleware.RolesAndPermissionMiddleware(handler.CreatePoll()))
	pollHandler.Get("list/", middleware.RolesAndPermissionMiddleware(handler.ListPoll()))
	pollHandler.Patch("update/:id/", middleware.RolesAndPermissionMiddleware(handler.UpdatePoll()))
	pollHandler.Delete("delete/:id/", middleware.RolesAndPermissionMiddleware(handler.DeletePoll()))
	pollHandler.Patch("vote/:id/", middleware.RolesAndPermissionMiddleware(handler.PollVote()))

}
