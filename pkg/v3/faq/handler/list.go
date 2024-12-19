package handler

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListFaq() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := presenter.FAQListReq{
			Page:     int(utils.CheckPageInQuery(c)),
			PageSize: int(utils.CheckPageSizeInQuery(c)),
		}

		faq, totalPage, err := handler.usecase.ListFaq(request)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if faq == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        faq,
		}
		if request.Page > int(totalPage) {
			response.CurrentPage = int32(totalPage)
		}
		return c.JSON(response)
	}
}
