package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) AssignSubjectToTeacher() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := uint(utils.StringToUint(c.Params("userID")))
		type SubjectIDs struct { // had to do this as the request body was not being parsed properly, it should have been but it wasn't and i don't know why
			SubjectIDs []uint `json:"subjectIDs"`
		}
		// Parse the request body to retrieve the subject IDs
		var subjects SubjectIDs
		if err := c.BodyParser(&subjects); err != nil {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(map[string]string{"error": err.Error()}))
		}

		errMap := handler.usecase.AssignSubjectsToTeacher(userID, subjects.SubjectIDs)

		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
