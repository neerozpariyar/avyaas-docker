package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) GetContentDetails() fiber.Handler {
	return func(c *fiber.Ctx) error {
		contentID := uint(utils.StringToUint(c.Params("id")))
		requesterID := c.Locals("requester").(uint)

		content, errMap := handler.usecase.GetContentDetails(contentID, requesterID)
		if len(errMap) != 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if content == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.DetailResponse{
			Success: true,
			Data:    content,
		}

		return c.JSON(response)
	}
}
