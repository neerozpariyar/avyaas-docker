package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListFeedback() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		page := utils.CheckPageInQuery(c)
		pageSize := utils.CheckPageSizeInQuery(c)
		courseID := uint(utils.StringToUint(c.Query("courseID")))

		feedbacks, totalPage, err := handler.usecase.ListFeedback(page, courseID, pageSize)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if feedbacks == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(page),
			TotalPage:   int32(totalPage),
			Data:        feedbacks,
		}

		if int32(page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
