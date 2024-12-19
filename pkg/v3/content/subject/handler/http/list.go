package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		page := utils.CheckPageInQuery(c)
		pageSize := utils.CheckPageSizeInQuery(c)

		courseID := uint(utils.StringToUint(c.Query("courseID")))

		search := c.Query("search")

		subjects, totalPage, err := handler.usecase.ListSubject(page, courseID, search, pageSize)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if subjects == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(page),
			TotalPage:   int32(totalPage),
			Data:        subjects,
		}

		if int32(page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
