package http

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) AddFCMToken() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var fcmToken models.FCMToken
		errMap := make(map[string]string)

		userID := c.Locals("requester").(uint)
		if err := c.BodyParser(&fcmToken); err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		fcmToken.UserID = userID
		if err := handler.usecase.AddFCMToken(fcmToken); err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		if err := handler.usecase.FCMRegister(fcmToken); err != nil {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		return c.JSON(presenter.SuccessResponse())
	}
}
