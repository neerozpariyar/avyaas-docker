package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
RegisterStudent is a Fiber handler function for user registration. It parses the request JSON body,
validates the user information, and invokes the use case to register the user. If any validation
or registration error occurs, it returns an appropriate error response. Otherwise, it returns a
success response
*/
func (handler *handler) RegisterStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.StudentRegisterRequest

		// Parse the request json body to validate and extract the user information
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		errMap2 := make(map[string]string)

		// Check if both new password matches
		if requestBody.Password != requestBody.ConfirmPassword {
			errMap2["password"] = "passwords do not match"
			errMap2["confirmPassword"] = "passwords do not match"
		}

		// Validates the request validateBody using a validator and returns validation errors if present
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

		// Convert the request data to JSON data
		// data, err := json.Marshal(requestBody)
		// if err != nil {
		// 	errMap["error"] = err.Error()
		// 	return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		// }

		// // Convert JSON data to a models.User instance
		// err = json.Unmarshal(data, &user)
		// if err != nil {
		// 	errMap["error"] = err.Error()
		// 	return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		// }

		// Invoke the registration of user
		// errMap = handler.usecase.RegisterStudent(*user)
		// if len(errMap) > 0 {
		// 	return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		// }
		requestBody.Referral = c.Query("referral")

		errMap = handler.usecase.RegisterStudent(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
