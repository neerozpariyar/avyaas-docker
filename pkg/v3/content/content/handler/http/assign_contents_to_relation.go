package http

import (
	"avyaas/internal/domain/presenter"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) AssignContentsToRelation() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody presenter.AssignContentsToRelation

		errMap := make(map[string]string)

		err := c.BodyParser(&requestBody)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		errMap = handler.usecase.AssignContentsToRelation(requestBody.RelationID, requestBody.SubjectID, requestBody.ContentIDs)

		if len(errMap) != 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())

	}
}
