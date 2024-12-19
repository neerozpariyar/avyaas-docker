package utils

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

/*
CheckPageInQuery is a utility function for parsing and validating the "page" query parameter from a
Fiber context. It takes a Fiber context as input, extracts the "page" query parameter, and returns
the parsed integer value. If the "page" parameter is not provided or is less than 1, the function
defaults to returning 1.

Parameters:
  - c: A pointer to the Fiber context representing the HTTP request context.

Returns:
  - int: The parsed and validated page number. Defaults to 1 if the "page" parameter is not provided
    or is less than 1.
*/
func CheckPageInQuery(c *fiber.Ctx) int {
	// Set default page value to 1
	page := 1

	// Check if page value is set in query params. If yes, convert it to int
	if c.Query("page") != "" {
		// Parse the "page" query parameter or set it to 1 if not provided or less than 0
		pageRequest, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			log.Println(err)
		}

		if pageRequest <= 0 {
			page = 1
		} else {
			page = pageRequest
		}
	}

	return page
}

func CheckPageSizeInQuery(c *fiber.Ctx) int {
	var err error
	pageRequest := 0
	pageSize := 0

	// Check if page value is set in query params. If yes, convert it to int
	if c.Query("pageSize") != "" {
		// Parse the "page" query parameter or set it to 1 if not provided or less than 0
		pageRequest, err = strconv.Atoi(c.Query("pageSize"))
		if err != nil {
			log.Println(err)
		}
	}

	if pageRequest <= 0 {
		pageSize = viper.GetInt("pagination.page_size")
	} else {
		pageSize = pageRequest
	}

	return pageSize
}
