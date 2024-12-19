package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListTeacherReferrals() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		id := utils.StringToUint(c.Params("id"))

		referrals, err := handler.usecase.ListTeacherReferrals(uint(id))

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if referrals == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}

		return c.JSON(presenter.ListResponse{Data: referrals})
	}
}
