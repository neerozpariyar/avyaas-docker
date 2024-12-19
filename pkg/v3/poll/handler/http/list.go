package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListPoll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := &presenter.PollListRequest{
			Page:      utils.CheckPageInQuery(c),
			PageSize:  utils.CheckPageSizeInQuery(c),
			UserID:    c.Locals("requester").(uint),
			SubjectID: uint(utils.StringToUint(c.Query("subjectID"))),
			Search:    c.Query("search"),
		}

		contents, totalPage, err := handler.usecase.ListPoll(*request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if contents == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        contents,
		}

		if int32(request.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}
		return c.JSON(response)
	}
}
