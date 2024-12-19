package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
ListQuestion is a Fiber handler function that handles the list question endpoint. It queries the
question service to retrieve a paginated list of the questions.

Parameters:
  - c: The Fiber Context representing the HTTP request and response.

Returns:
  - error: An error, if any, encountered during the handling of the HTTP request.
*/
func (handler *handler) ListTypeQuestion() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := presenter.ListQuestionRequest{
			Page:          utils.CheckPageInQuery(c),
			PageSize:      utils.CheckPageSizeInQuery(c),
			CourseID:      uint(utils.StringToUint(c.Query("courseID"))),
			SubjectID:     uint(utils.StringToUint(c.Query("subjectID"))),
			QuestionSetID: uint(utils.StringToUint(c.Query("questionSetID"))),
			RequesterID:   c.Locals("requester").(uint),
			Search:        c.Query("search"),
		}

		// Invoke the test list usecase to retrieve the list of tests
		questions, totalPage, err := handler.usecase.ListTypeQuestion(request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if questions == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		// Initialize ListResponse presenter
		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        questions,
		}

		// Set the currentPage to value of totalPage if requested page is greater that totalPage
		if int32(request.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
