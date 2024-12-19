package http

import (
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) GetTestLeaderboard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		testID := uint(utils.StringToUint(c.Params("id")))

		response, err := handler.usecase.GetTestLeaderboard(testID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"data": response})
	}
}
