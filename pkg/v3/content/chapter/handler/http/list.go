package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListChapter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		data := presenter.ChapterListRequest{
			Page:     utils.CheckPageInQuery(c),
			PageSize: utils.CheckPageSizeInQuery(c),

			ChapterFilter: presenter.FilterChapterRequest{
				UnitID:    uint(utils.StringToUint(c.Query("unitID"))),
				SubjectID: uint(utils.StringToUint(c.Query("subjectID"))),
			},

			Search: c.Query("search"),
		}

		if data.ChapterFilter.UnitID != 0 || data.ChapterFilter.SubjectID != 0 {
			validate, trans := utils.InitTranslator()

			err := validate.Struct(data.ChapterFilter)
			if err != nil {
				validationErrors := err.(validator.ValidationErrors)
				errMap = utils.TranslateError(validationErrors, trans)

				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}
		}

		chapters, totalPage, err := handler.usecase.ListChapter(data)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if chapters == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		response := presenter.ChapterListResponse{
			Success:     true,
			CurrentPage: int32(data.Page),
			TotalPage:   int32(totalPage),
			Data:        chapters,
		}

		if int32(data.Page) > int32(totalPage) {
			response.CurrentPage = int32(totalPage)
		}

		return c.JSON(response)
	}
}
