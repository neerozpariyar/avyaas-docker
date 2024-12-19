package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) DeleteLive() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		
		// Convert and set the ID from the route parameter
		id := utils.StringToUint(c.Params("id"))

		// Invoke the live  update request to the live  usecase
		err := handler.usecase.DeleteLive(uint(id))
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
