package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListService() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		page := utils.CheckPageInQuery(c)
		search := c.Query("search")
		pageSize := utils.CheckPageSizeInQuery(c)
		services, totalPage, err := handler.usecase.ListService(page, search, pageSize)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if len(services) == 0 {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		// Initialize ListResponse presenter
		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(page),
			TotalPage:   int32(totalPage),
			Data:        services,
		}

		// Set the currentPage to value of totalPage if requested page is greater that totalPage
		if int32(page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
