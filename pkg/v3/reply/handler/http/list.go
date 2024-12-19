package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListReply() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := &presenter.ReplyListRequest{
			Page:         utils.CheckPageInQuery(c),
			PageSize:     utils.CheckPageSizeInQuery(c),
			DiscussionID: uint(utils.StringToUint(c.Query("discussionID"))),
			Search:       c.Query("search"),
		}

		replies, totalPage, err := handler.usecase.ListReply(*request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if replies == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        replies,
		}

		if int32(request.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
