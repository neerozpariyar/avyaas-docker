package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) MeetingSDKKey() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		var requestBody struct {
			AppKey    string `json:"appKey"`
			SecretKey string `json:"secretKey"`
			MeetingID int64  `json:"meetingID"`
			Role      int    `json:"role"`
		}
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

		msdk, err := handler.usecase.MeetingSDKKey(requestBody.AppKey, requestBody.MeetingID, requestBody.Role, requestBody.SecretKey)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		response := presenter.ListResponse{
			Success: true,
			Data:    msdk,
		}

		return c.JSON(response)
	}
}
