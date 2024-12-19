package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"strings"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CreateReferral() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.CreateUpdateReferralRequest

		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()

		errMap2 := make(map[string]string)
		if *requestBody.HasLimit && requestBody.Limit == 0 {
			errMap2["limit"] = "limit is a required field"
		}

		if strings.ToUpper(requestBody.Type) == "COURSE" && requestBody.CourseID == 0 {
			errMap2["courseID"] = "courseID is a required field"
		}

		if strings.ToUpper(requestBody.Type) != "GENERAL" && strings.ToUpper(requestBody.Type) != "COURSE" {
			errMap2["type"] = "invalid type value"
		}

		requestBody.Type = strings.ToUpper(requestBody.Type)
		requestBody.DiscountType = strings.ToUpper(requestBody.DiscountType)

		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			if len(errMap2) != 0 {
				errMap = utils.MergeMaps(errMap, errMap2)
			}
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		if len(errMap2) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap2))
		}

		// Invoke request to the content usecase to create the content
		errMap = handler.usecase.CreateReferral(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
