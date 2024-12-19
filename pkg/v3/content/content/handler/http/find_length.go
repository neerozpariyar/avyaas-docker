package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FileLengthRequest struct {
	ObjectKey string `json:"objectKey"`
}

func (handler *handler) FindVideoLength() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		// Get the file from file field
		// file, _ := c.FormFile("file")
		// if file == nil {
		// 	errMap["file"] = "file is a required field"
		// 	return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		// }
		// abc := file.Header
		// fmt.Printf("abc: %v\n", abc)
		// fmt.Printf("abc.Get(\"length\"): %v\n", abc.Get("Content-Length"))

		// rData, err := file.Open()
		// if err != nil {
		// 	errMap["error"] = err.Error()
		// 	return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		// }

		// Note: change the fle length check package later
		// pData, err := ffprobe.ProbeReader(context.Background(), rData)

		var requestBody FileLengthRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()

		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		response, err := file.GetFileLength(requestBody.ObjectKey)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(response)
	}
}
