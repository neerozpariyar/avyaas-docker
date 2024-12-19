package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListTestSeries() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := presenter.ListTestSeriesRequest{
			Page:     utils.CheckPageInQuery(c),
			PageSize: utils.CheckPageSizeInQuery(c),
			Search:   c.Query("search"),
			CourseID: uint(utils.StringToUint(c.Query("courseID"))),
		}

		// Invoke the test series list usecase to retrieve the list of test series
		testSeries, totalPage, err := handler.usecase.ListTestSeries(request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if len(testSeries) == 0 {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        testSeries,
		}

		// Set the currentPage to value of totalPage if requested page is greater that totalPage
		if int32(request.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
