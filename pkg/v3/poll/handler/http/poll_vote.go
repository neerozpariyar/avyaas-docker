package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) PollVote() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		var requestBody presenter.PollVoteRequest
		if err := c.BodyParser(&requestBody); err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		requestBody.PollID = uint((utils.StringToUint(c.Params("id"))))
		requestBody.UserID = c.Locals("requester").(uint)

		errMap = handler.usecase.PollVote(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))

		}

		return c.JSON(presenter.SuccessResponse())
	}
}
