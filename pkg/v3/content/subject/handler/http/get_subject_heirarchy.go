package http

import (
	"avyaas/internal/domain/presenter"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) GetSubjectHeirarchy() fiber.Handler {
	return func(c *fiber.Ctx) error {

		errMap := make(map[string]string)

		id, err := strconv.Atoi(c.Params("id"))

		userID := c.Locals("requester").(uint)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		result, err := handler.usecase.GetSubjectHeirarchy(uint(id), userID)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(fiber.Map{"data": result})

	}
}
