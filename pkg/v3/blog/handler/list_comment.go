package handler

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListComments() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := presenter.BlogCommentListReq{
			Page:     int(utils.CheckPageInQuery(c)),
			PageSize: int(utils.CheckPageSizeInQuery(c)),
			Search:   c.Query("search"),
			UserID:   uint(utils.StringToUint(c.Query("userID"))),
			BlogID:   uint(utils.StringToUint(c.Query("blogID"))),
		}

		comment, totalPage, err := handler.usecase.ListComments(request)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if comment == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        comment,
		}
		if request.Page > totalPage {
			response.CurrentPage = int32(totalPage)
		}
		return c.JSON(response)
	}
}
