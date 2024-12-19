package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) DeleteReferral() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		id := utils.StringToUint(c.Params("id"))

		err := handler.usecase.DeleteReferral(uint(id))
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
