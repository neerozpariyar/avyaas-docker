package http

import (
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) GetTestResult() fiber.Handler {
	return func(c *fiber.Ctx) error {
		testID := uint(utils.StringToUint(c.Params("id")))
		requesterID := c.Locals("requester").(uint)

		if testID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "testID is required"})
		}

		result, err := handler.usecase.GetTestResult(testID, requesterID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"data": result})
	}
}
