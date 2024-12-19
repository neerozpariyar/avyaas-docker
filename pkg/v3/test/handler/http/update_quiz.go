package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"time"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
CreateTest handles the HTTP request to create a new test by sending a HTTP request to the test
usecase.

Parameters:
  - c: The Fiber Context representing the HTTP request and response.

Returns:
  - error: An error, if any, encountered during the handling of the HTTP request.
*/
func (handler *handler) UpdateTest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.CreateUpdateTestRequest

		// Parse the request json body to validate and extract the test information
		if err := c.BodyParser(&requestBody); err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		errMap2 := make(map[string]string)
		if !requestBody.IsPremium && requestBody.TestSeriesID != 0 {
			errMap2["isPremium"] = "isPremium is a required field"
		}

		if requestBody.IsPremium && requestBody.TestSeriesID == 0 {
			errMap2["testSeriesID"] = "testSeriesID is a required field"
		}

		// Check if test is premium and if price is given
		if !requestBody.IsFree && requestBody.Price <= 0 {
			if !requestBody.IsPremium {
				errMap2["price"] = "price is a required field"
			}
		}

		if requestBody.IsPublic && requestBody.QuestionSetID == 0 {
			errMap2["isDraft"] = "test cannot be made public without a question set assigned"
		}

		if st, err := utils.ParseStringToTime(requestBody.StartTime); err == nil {
			if st.Unix() < time.Now().Unix() {
				errMap2["startTime"] = "startTime cannot be earlier to current datetime"
			}
		}

		if et, err := utils.ParseStringToTime(requestBody.EndTime); err == nil {
			if et.Unix() < time.Now().Unix() {
				errMap2["endTime"] = "endTime cannot be earlier to current datetime"
			}
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

		// Get the test if from params
		requestBody.ID = uint(utils.StringToUint(c.Params("id")))

		// Invoke request to the test usecase to create the test
		if errMap = handler.usecase.UpdateTest(requestBody); len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
