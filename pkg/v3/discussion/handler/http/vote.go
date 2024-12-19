package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) LikeOrUnlikeDiscussion() fiber.Handler {
	return func(c *fiber.Ctx) error {
		discussionID := uint((utils.StringToUint(c.Params("id"))))
		var requestBody presenter.DiscussionCreateUpdateRequest

		requestBody.UserID = c.Locals("requester").(uint)

		errMap := handler.usecase.LikeOrUnlikeDiscussion(discussionID, requestBody.UserID)
		if len(errMap) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		return c.JSON(presenter.SuccessResponse())
	}
}
