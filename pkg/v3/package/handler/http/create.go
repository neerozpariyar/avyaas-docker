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
func (handler *handler) CreatePackage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var requestBody presenter.PackageCreateUpdateRequest

		// Parse the request json body to validate and extract the package information
		if err := c.BodyParser(&requestBody); err != nil {
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

		// packageType := requestBody.PackageTypeID
		// switch packageType {
		// // NOTE: This might need to be changed based on package types required
		// case 1, 4:
		// 	if requestBody.TestSeriesID == 0 {
		// 		errMap2["testSeriesID"] = "testSeriesID is a required field"
		// 	}

		// 	if requestBody.LiveGroupID == 0 {
		// 		errMap2["liveGroupID"] = "liveGroupID is a required field"
		// 	}
		// case 5:
		// 	if requestBody.CourseID == 0 {
		// 		errMap2["courseID"] = "courseID is a required field"
		// 	}
		// case 2, 6:
		// 	if requestBody.TestSeriesID == 0 {
		// 		errMap2["testSeriesID"] = "testSeriesID is a required field"
		// 	}
		// case 3, 7:
		// 	if requestBody.LiveGroupID == 0 {
		// 		errMap2["liveGroupID"] = "liveGroupID is a required field"
		// 	}
		// case 8:
		// 	if requestBody.TestID == 0 {
		// 		errMap2["testID"] = "testID is a required field"
		// 	}
		// default:
		// 	errMap2["error"] = "invalid package type"
		// }

		// Validates the request validateBody using a validator and returns validation errors if present
		if err := validate.Struct(requestBody); err != nil {
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

		// Invoke request to the package usecase to create the test
		if errMap = handler.usecase.CreatePackage(requestBody); len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}
}
