package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) InitiateKhaltiPayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		var request presenter.InitateKhaltiPaymentRequest
		request.UserID = c.Locals("requester").(uint)

		err := c.BodyParser(&request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		validate, trans := utils.InitTranslator()

		err = validate.Struct(request)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		response, errMap := handler.usecase.InitiateKhaltiPayment(request)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.Status(http.StatusOK).JSON(presenter.DetailResponse{Success: true, Data: response})
	}
}
