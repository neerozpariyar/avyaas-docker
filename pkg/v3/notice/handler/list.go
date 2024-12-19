package handler

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListNotice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := &presenter.NoticeListReq{
			Page:     utils.CheckPageInQuery(c),
			PageSize: utils.CheckPageSizeInQuery(c),
			CourseID: uint(utils.StringToUint(c.Query("courseID"))),
			Search:   c.Query("search"),
		}

		notices, totalPage, err := handler.usecase.ListNotice(*request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if notices == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        notices,
		}

		if int32(request.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
