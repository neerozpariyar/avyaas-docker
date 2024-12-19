package http

import (
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) GetQuestionSetDetails() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		id := uint(utils.StringToUint(c.Params("id")))
		userID := c.Locals("requester").(uint)

		// Invoke the test list usecase to retrieve the list of question sets
		response, err := handler.usecase.GetQuestionSetDetails(id, userID)
		if err != nil {
			errMap["errors"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(errMap)
		}

		return c.JSON(fiber.Map{"data": response})
	}
}
