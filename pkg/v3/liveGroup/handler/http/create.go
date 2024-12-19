package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CreateLiveGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		// var requestBody models.LiveGroup
		var requestBody presenter.LiveGroupCreateUpdatePresenter

		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()

		errMap2 := make(map[string]string)

		if requestBody.IsPackage && requestBody.Price == 0 {
			errMap2["price"] = "price is a required field"
		}

		if requestBody.IsPackage && requestBody.Period == 0 {
			errMap2["period"] = "period is a required field"
		}

		if requestBody.IsPackage && requestBody.PackageTypeID == 0 {
			errMap2["packageTypeID"] = "packageTypeID is a required field"
		}

		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)

			errMap = utils.TranslateError(validationErrors, trans)
			if len(errMap2) != 0 {
				errMap = utils.MergeMaps(errMap, errMap2)
			}

		}

		if len(errMap2) != 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap2))
		}

		newErrMap := handler.usecase.CreateLiveGroup(requestBody)
		errMap = utils.MergeMaps(errMap, newErrMap)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
