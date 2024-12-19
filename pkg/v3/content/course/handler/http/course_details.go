package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) GetCourseDetails() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		courseID := uint(utils.StringToUint(c.Params("id")))
		userID := c.Locals("requester").(uint)

		courseDetails, err := handler.usecase.GetCourseDetails(userID, courseID)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(fiber.Map{"data": courseDetails})
	}
}
