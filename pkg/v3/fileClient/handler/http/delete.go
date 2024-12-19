package http

import (
	"avyaas/internal/domain/presenter"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) DeleteObjects() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		// Convert and set the ID from the route parameter
		var ids presenter.DeleteObjectReq
		if err := c.BodyParser(&ids); err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		for _, id := range ids.Ids {
			err := handler.usecase.DeleteObjects([]uint{id})
			if err != nil {
				errMap["error"] = err.Error()
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}
		}

		return c.JSON(presenter.SuccessResponse())

	}
}
