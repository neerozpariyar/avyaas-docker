package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListObjects() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		req := &presenter.FileListReq{}
		req.Page = utils.CheckPageInQuery(c)
		req.PageSize = utils.CheckPageSizeInQuery(c)
		req.Service = c.Query("service")
		req.Search = c.Query("search")
		isActiveStr := c.Query("isActive")
		if isActiveStr == "" {
			// Set a default value for isActive if it is not provided
			req.IsActive = true
		} else {
			isActive, err := strconv.ParseBool(isActiveStr) // converts the string to a boolean
			if err != nil {
				errMap["error"] = err.Error()
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}
			req.IsActive = isActive
		}
		files, totalPage, err := handler.usecase.ListObjects(req)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if files == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(req.Page),
			TotalPage:   int32(totalPage),
			Data:        files,
		}

		if int32(req.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
