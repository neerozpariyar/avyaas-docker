package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
ListTeacher is a Fiber handler function for processing requests to retrieve a paginated list of teachers.
It accepts optional query parameters to specify the page number and invokes the teacher list use case.

Parameters:
  - handler: A pointer to the handler struct, typically representing the application's request handler.
    It is used to access the associated use case for listing teachers.

Returns:
  - fiber.Handler: A Fiber handler function that can be registered to handle HTTP requests. The handler
    function processes requests, performs pagination-related checks, and delegates the teacher list
    retrieval operation to the use case.
*/
func (handler *handler) ListTeacher() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := &presenter.TeacherListRequest{
			Page:      utils.CheckPageInQuery(c),
			PageSize:  utils.CheckPageSizeInQuery(c),
			CourseID:  uint((utils.StringToUint(c.Query("courseID")))),
			SubjectID: uint((utils.StringToUint(c.Query("subjectID")))),
			Search:    c.Query("search"),
		}

		// Invoke the teacher list usecase to retrieve the list of teacher
		teachers, totalPage, err := handler.usecase.ListTeacher(request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if teachers == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		// Initialize TeacherListResponse presenter
		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        teachers,
		}

		// Set the currentPage to value of totalPage if requested page is greater that totalPage
		if int32(request.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
