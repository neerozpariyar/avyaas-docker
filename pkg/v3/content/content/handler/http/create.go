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

func (handler *handler) CreateContent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.ContentCreateUpdateRequest

		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()

		errMap2 := make(map[string]string)
		if strings.ToUpper(requestBody.ContentType) == "VIDEO" && requestBody.Length == 0 {
			errMap2["length"] = "length is a required field"
		}

		if requestBody != (presenter.ContentCreateUpdateRequest{}) && requestBody.HasNote != nil && *requestBody.HasNote {
			if requestBody.Note.Title == "" {
				errMap2["note.title"] = "title is a required field"
			}

			noteFile, err := c.FormFile("note.file")
			if err != nil {
				requestBody.Note.File = nil
			} else {
				requestBody.Note.File = noteFile
			}
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

		requestBody.CreatedBy = c.Locals("requester").(uint)

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
		// Invoke request to the content usecase to create the content
		errMap = handler.usecase.CreateContent(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
