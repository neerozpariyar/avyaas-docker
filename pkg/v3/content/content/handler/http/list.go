package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListContent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := &presenter.ContentListRequest{
			Page:     utils.CheckPageInQuery(c),
			PageSize: utils.CheckPageSizeInQuery(c),
			ContentFilter: presenter.FilterContentRequest{
				ChapterID: uint(utils.StringToUint(c.Query("chapterID"))),
				SubjectID: uint(utils.StringToUint(c.Query("subjectID"))),
				UnitID:    uint(utils.StringToUint(c.Query("unitID"))),
			},
			Search: c.Query("search"),
			UserID: c.Locals("requester").(uint),
		}

		if request.ContentFilter.ChapterID != 0 || request.ContentFilter.SubjectID != 0 || request.ContentFilter.UnitID != 0 {

			validate, trans := utils.InitTranslator()

			err := validate.Struct(request.ContentFilter)
			if err != nil {
				validationErrors := err.(validator.ValidationErrors)
				errMap = utils.TranslateError(validationErrors, trans)

				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}
		}

		contents, totalPage, err := handler.usecase.ListContent(*request)
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
