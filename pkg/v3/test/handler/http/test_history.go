package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) GetTestHistory() fiber.Handler {
	errMap := make(map[string]string)

	return func(c *fiber.Ctx) error {
		request := presenter.TestHistoryRequest{
			UserID:   c.Locals("requester").(uint),
			CourseID: uint(utils.StringToUint(c.Query("courseID"))),
			FromDate: c.Query("fromDate"),
			ToDate:   c.Query("toDate"),
			Page:     utils.CheckPageInQuery(c),
			PageSize: utils.CheckPageSizeInQuery(c),
		}

		data, totalPage, err := handler.usecase.GetTestHistory(request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize ListResponse presenter
		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        data,
		}

		// Set the currentPage to value of totalPage if requested page is greater that totalPage
		if int32(request.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
