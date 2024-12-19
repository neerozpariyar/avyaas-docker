package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListNote() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		data := presenter.NoteListRequest{
			Page:      utils.CheckPageInQuery(c),
			PageSize:  utils.CheckPageSizeInQuery(c),
			ContentID: uint(utils.StringToUint(c.Query("contentID"))),
			Search:    c.Query("search"),
		}

		notes, totalPage, err := handler.usecase.ListNote(data)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if notes == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(data.Page),
			TotalPage:   int32(totalPage),
			Data:        notes,
		}

		if int32(data.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
