package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) GetBookmarkDetails() fiber.Handler {
	return func(c *fiber.Ctx) error {
		bookmarkID := uint(utils.StringToUint(c.Params("id")))
		// requesterID := c.Locals("requester").(uint)

		bookmark, errMap := handler.usecase.GetBookmarkDetails(bookmarkID)
		if len(errMap) != 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if bookmark == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.DetailResponse{
			Success: true,
			Data:    bookmark,
		}

		return c.JSON(response)
	}
}
