package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// here if courseGroupID available, generates dynamic list using query params otherwise displays all Courses
func (handler *handler) ListCourse() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := presenter.CourseListRequest{
			UserID:        c.Locals("requester").(uint),
			CourseGroupID: uint((utils.StringToUint(c.Query("courseGroupID")))),
			Search:        c.Query("search"),
			Page:          utils.CheckPageInQuery(c),
			PageSize:      utils.CheckPageSizeInQuery(c),
		}
		courses, totalPage, err := handler.usecase.ListCourse(request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if courses == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        courses,
		}

		if int32(request.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
