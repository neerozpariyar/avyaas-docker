package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) UpdateTestSeries() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.TestSeriesCreateUpdateRequest

		// Parse and validate the request json body
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		errMap2 := make(map[string]string)

		// Validate if the required data are provided to create test series as a package
		if requestBody.IsPackage {
			errMap2 = ValidateTestSeriesPackageData(requestBody)
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

		// Convert and set the ID from the route parameter to requestBody
		requestBody.ID = uint(utils.StringToUint(c.Params("id")))

		// Invoke request to the test series usecase to update the test series
		errMap = handler.usecase.UpdateTestSeries(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
