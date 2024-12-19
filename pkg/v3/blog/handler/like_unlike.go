package handler

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) BlogLikeUnlike() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		uID := c.Locals("requester").(uint)
		bID := utils.StringToUint(c.Params("id"))

		_, err := handler.usecase.BlogLikeUnlike(uID, uint(bID))
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		return c.JSON(presenter.SuccessResponse())
	}
}
