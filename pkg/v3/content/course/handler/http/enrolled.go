package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// here if courseGroupID available, generates dynamic list using query params otherwise displays all Courses
func (handler *handler) ListEnrolledCourse() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		page := utils.CheckPageInQuery(c)
		search := c.Query("search")
		userID := c.Locals("requester").(uint)
		pageSize := utils.CheckPageSizeInQuery(c)

		courses, totalPage, err := handler.usecase.ListEnrolledCourse(userID, page, search, pageSize)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if len(courses) == 0 {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(page),
			TotalPage:   int32(totalPage),
			Data:        courses,
		}

		if int32(page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
