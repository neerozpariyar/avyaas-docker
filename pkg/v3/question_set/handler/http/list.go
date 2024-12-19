package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
ListQuestionSet is a Fiber handler function that handles the list question set endpoint. It queries
the question set service to retrieve a paginated list of the question sets.

Parameters:
  - c: The Fiber Context representing the HTTP request and response.

Returns:
  - error: An error, if any, encountered during the handling of the HTTP request.
*/
func (handler *handler) ListQuestionSet() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		// Check if page value is set in query params
		page := utils.CheckPageInQuery(c)
		pageSize := utils.CheckPageSizeInQuery(c)
		// Get the courseID value from query params if set
		courseID := uint(utils.StringToUint(c.Query("courseID")))

		// Invoke the test list usecase to retrieve the list of question sets
		questionSets, totalPage, err := handler.usecase.ListQuestionSet(page, courseID, pageSize)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if questionSets == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		// Initialize ListResponse presenter
		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(page),
			TotalPage:   int32(totalPage),
			Data:        questionSets,
		}

		// Set the currentPage to value of totalPage if requested page is greater that totalPage
		if int32(page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
