package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) DeleteLiveGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		// Convert and set the ID from the route parameter
		id := utils.StringToUint(c.Params("id"))

		// Invoke the liveGroup  update request to the liveGroup  usecase
		err := handler.usecase.DeleteLiveGroup(uint(id))
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
