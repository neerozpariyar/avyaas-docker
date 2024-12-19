package handler

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListBlog() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := presenter.BlogListReq{
			Page:      int(utils.CheckPageInQuery(c)),
			PageSize:  int(utils.CheckPageSizeInQuery(c)),
			CourseID:  uint(utils.StringToUint(c.Query("courseID"))),
			SubjectID: uint(utils.StringToUint(c.Query("subjectID"))),
			Search:    c.Query("search"),
		}

		blog, totalPage, err := handler.usecase.ListBlog(request)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if blog == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		response := presenter.ListResponse{
			Success:     true,
			CurrentPage: int32(request.Page),
			TotalPage:   int32(totalPage),
			Data:        blog,
		}
		if request.Page > totalPage {
			response.CurrentPage = int32(totalPage)
		}
		return c.JSON(response)
	}
}
