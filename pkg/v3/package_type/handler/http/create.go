package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
CreatePackage handles the HTTP request to create a new package by sending a HTTP request to the package
usecase.

Parameters:
  - c: The Fiber Context representing the HTTP request and response.

Returns:
  - error: An error, if any, encountered during the handling of the HTTP request.
*/
func (handler *handler) CreatePackageType() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.PackageTypeCreateUpdateRequest

		// Parse the request json body to validate and extract the package information
		if err := c.BodyParser(&requestBody); err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		errMap2 := make(map[string]string)
		if utils.ContainsUint(requestBody.ServiceIDs, 2) && utils.ContainsUint(requestBody.ServiceIDs, 3) {
			errMap2["testServiceIDs"] = "package type cannot have both test series and test"
		}

		if utils.ContainsUint(requestBody.ServiceIDs, 4) && utils.ContainsUint(requestBody.ServiceIDs, 5) {
			errMap2["liveServiceIDs"] = "package type cannot have both live group and live"
		}

		// Validates the request validateBody using a validator and returns validation errors if present
		if err := validate.Struct(requestBody); err != nil {
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

		// Invoke request to the package usecase to create the test
		if errMap = handler.usecase.CreatePackageType(requestBody); len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
