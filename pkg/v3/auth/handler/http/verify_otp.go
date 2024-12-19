package http

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
VerifyUserOTP is a Fiber handler function that handles the verification of a user's OTP (One-Time
Password).
*/
func (handler *handler) VerifyUserOTP() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.VerifyOTPRequest
		var otpRequest *models.UserOtp

		// Parse the request json body to validate
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		// Validates the request validateBody using a validator and returns validation errors if present
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Convert the request data to JSON data
		data, err := json.Marshal(requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Convert JSON data to a models.UserOtp instance
		err = json.Unmarshal(data, &otpRequest)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Call the usecase for user otp verification
		verified, err := handler.usecase.VerifyUserOTP(*otpRequest)
		if verified {
			return c.JSON(presenter.SuccessResponse())
		}

		errMap["error"] = err.Error()
		return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))

	}
}
