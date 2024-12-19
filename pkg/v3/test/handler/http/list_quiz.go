package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
ListTest is a Fiber handler function that handles the list test endpoint. It queries the test service
to retrieve a paginated list of the tests.

Parameters:
  - c: The Fiber Context representing the HTTP request and response.

Returns:
  - error: An error, if any, encountered during the handling of the HTTP request.
*/
func (handler *handler) ListTest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := presenter.ListTestRequest{
			UserID:     c.Locals("requester").(uint),
			CourseID:   uint(utils.StringToUint(c.Query("courseID"))),
			Page:       utils.CheckPageInQuery(c),
			PageSize:   utils.CheckPageSizeInQuery(c),
			TestTypeID: uint(utils.StringToUint(c.Query("testTypeID"))),
			Status:     c.Query("status"),
			FromDate:   c.Query("fromDate"),
			ToDate:     c.Query("toDate"),
		}

		// Invoke the test list usecase to retrieve the list of tests
		tests, totalPage, err := handler.usecase.ListTest(request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if len(tests) == 0 {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		// Initialize ListResponse presenter
		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        tests,
		}

		// Set the currentPage to value of totalPage if requested page is greater that totalPage
		if int32(request.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
