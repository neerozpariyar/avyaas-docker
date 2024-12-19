package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) UpdateContent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.ContentCreateUpdateRequest

		// Parse and validate the request json body
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		errMap2 := make(map[string]string)
		if strings.ToUpper(requestBody.ContentType) == "VIDEO" && requestBody.Length == 0 {
			errMap2["length"] = "length is a required field"
		}

		// Validate the request validateBody using a validator and return the translated validation errors if present
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			if len(errMap2) != 0 {
				errMap = utils.MergeMaps(errMap, errMap2)
			}

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		if len(errMap2) != 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap2))
		}

		// Get the file from file field
		file, _ := c.FormFile("file")
		if file != nil {
			fileType := utils.GetFileType(file.Filename)
			if strings.ToUpper(requestBody.ContentType) == "VIDEO" {
				if fileType != "mpeg" && fileType != "mp4" && fileType != "mov" { //validate file type before setting
					errMap["file"] = fmt.Errorf("file type of %v not allowed for video content: only video type: mpeg. mp4 and mov allowed", fileType).Error()
					return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
				}
			} else if strings.ToUpper(requestBody.ContentType) == "PDF" {
				if fileType != "pdf" { //validate file type before setting
					errMap["file"] = fmt.Errorf("file type of %v not allowed for resourse content: only file type: pdf allowed", fileType).Error()
					return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
				}
			} else {
				errMap["contentType"] = fmt.Errorf("invalid content type").Error()
				return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
			}

			requestBody.File = file
		} else {
			requestBody.File = nil
		}

		// Convert and set the ID from the route parameter to requestBody
		requestBody.ID = uint(utils.StringToUint(c.Params("id")))

		// Invoke request to the content  usecase to update the content
		errMap = handler.usecase.UpdateContent(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
