package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) UpdatePackageType() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.PackageTypeCreateUpdateRequest

		// Parse and validate the request json body
		err := c.BodyParser(&requestBody)
		if err != nil {
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

		// Invoke request to the package type usecase to update the package type
		errMap = handler.usecase.UpdatePackageType(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
