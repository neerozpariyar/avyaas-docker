package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CreateTestSeries() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.TestSeriesCreateUpdateRequest

		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()

		errMap2 := make(map[string]string)

		// Validate if the required data are provided to create test series as a package
		if requestBody.IsPackage {
			errMap2 = ValidateTestSeriesPackageData(requestBody)
		}

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

		// Invoke request to the chapter usecase to create the chapter
		newErrMap := handler.usecase.CreateTestSeries(requestBody)
		errMap = utils.MergeMaps(errMap, newErrMap)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}

func ValidateTestSeriesPackageData(request presenter.TestSeriesCreateUpdateRequest) map[string]string {
	errMap := make(map[string]string)

	if request.Price == 0 {
		errMap["price"] = "price is a required field"
	}

	if request.Period == 0 {
		errMap["period"] = "period is a required field"
	}

	if request.PackageTypeID == 0 {
		errMap["packageTypeID"] = "packageTypeID is a required field"
	}

	return errMap
}
