package http

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) UpdatePackage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.PackageCreateUpdateRequest

		// Parse and validate the request json body
		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		// Initialize the translation function
		validate, trans := utils.InitTranslator()

		errMap2 := make(map[string]string)

		serviceIDs, err := handler.packageTypeUsecase.GetPackageTypeServices(requestBody.PackageTypeID)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		for idx := range serviceIDs {
			switch serviceIDs[idx] {
			case 1:
				println("course")
				if requestBody.CourseID == 0 {
					errMap2["courseID"] = "courseID is a required field"
				}
			case 2:
				println("test series")
				if requestBody.TestSeriesID == 0 {
					errMap2["testSeriesID"] = "testSeriesID is a required field"
				}
			case 3:
				println("test")
				if requestBody.TestID == 0 {
					errMap2["testID"] = "testID is a required field"
				}
			case 4:
				println("live group")
				if requestBody.LiveGroupID == 0 {
					errMap2["liveGroupID"] = "liveGroupID is a required field"
				}
			case 5:
				println("live")
				if requestBody.LiveID == 0 {
					errMap2["liveID"] = "liveID is a required field"
				}
			default:
				errMap2["error"] = "invalid package type"
			}
		}

		// Validate the request validateBody using a validator and return the translated validation errors if present
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)

			if len(errMap2) != 0 {
				utils.MergeMaps(errMap, errMap2)
			}

			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		if len(errMap2) != 0 {
			utils.MergeMaps(errMap, errMap2)
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap2))
		}

		// Convert and set the ID from the route parameter to requestBody
		requestBody.ID = uint(utils.StringToUint(c.Params("id")))

		// Invoke request to the subject  usecase to update the subject
		errMap = handler.usecase.UpdatePackage(requestBody)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
