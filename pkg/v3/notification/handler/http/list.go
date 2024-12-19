package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListNotification() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		requestBody := presenter.NotificationListRequest{
			Page:     utils.CheckPageInQuery(c),
			PageSize: utils.CheckPageSizeInQuery(c),
			Search:   c.Query("search"),
			CourseID: uint(utils.StringToUint(c.Query("courseID"))),
		}

		notifications, totalPage, err := handler.usecase.ListNotification(requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if len(notifications) == 0 {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(requestBody.Page),
			TotalPage:   int32(totalPage),
			Data:        notifications,
		}

		if int32(requestBody.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
