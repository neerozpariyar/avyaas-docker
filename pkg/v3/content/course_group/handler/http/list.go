package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
ListCourseGroup is a Fiber handler function that handles the list course group endpoint. It queries
the course group service to retrieve a paginated list of course groups.

Parameters:
  - c: The Fiber Context representing the HTTP request and response.

Returns:
  - error: An error, if any, encountered during the handling of the HTTP request.
*/
func (handler *handler) ListCourseGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		// Check if page value is set in query params
		page := utils.CheckPageInQuery(c)
		pageSize := utils.CheckPageSizeInQuery(c)

		search := c.Query("search")

		// Invoke the course group list usecase to retrieve the list of course group
		courseGroups, totalPage, err := handler.usecase.ListCourseGroup(page, search, pageSize)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if len(courseGroups) == 0 {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		// Initialize CourseGroupListResponse presenter
		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(page),
			TotalPage:   int32(totalPage),
			Data:        courseGroups,
		}

		// Set the currentPage to value of totalPage if requested page is greater that totalPage
		if int32(page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
