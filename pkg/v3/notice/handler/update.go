package handler

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) UpdateNotice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var request presenter.NoticeCreateUpdatePresenter

		err := c.BodyParser(&request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()

		err = validate.Struct(request)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)
		}

		file, _ := c.FormFile("file")
		if file != nil {
			filteType := utils.GetFileType(file.Filename)

			if filteType != "png" && filteType != "jpg" && filteType != "jpeg" && filteType != "pdf" {
				errMap["file"] = fmt.Errorf("file type of %v not allowed : only png or jpg or jpeg or pdf allowed", filteType).Error()
			} else {
				errMap["file"] = fmt.Errorf("invalid file type").Error()
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}
			request.File = file
		}

		request.ID = uint(utils.StringToUint(c.Params("id")))

		_, errMap = handler.usecase.UpdateNotice(request)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		return c.JSON(presenter.SuccessResponse())
	}
}
