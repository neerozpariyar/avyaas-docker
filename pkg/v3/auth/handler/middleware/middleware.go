package middleware

import (
	"avyaas/internal/config"
	"avyaas/internal/core"
	"avyaas/internal/domain/presenter"
	accountRepo "avyaas/pkg/v3/account/repository/gorm"
	repository "avyaas/pkg/v3/auth/repository/gorm"
	"avyaas/pkg/v3/auth/usecase"
	courseRepo "avyaas/pkg/v3/content/course/repository/gorm"
	courseGroupRepo "avyaas/pkg/v3/content/course_group/repository/gorm"
	"fmt"
	"sync"

	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v2"
	"gorm.io/gorm"
)

/*
Protected is a Fiber middleware function that applies JWT protection to routes. It takes a secret
string as a parameter, which represents the key for validating JWT tokens. It uses the jwtWare
package to configure JWT protection with the provided secret, custom claims, and an error handler
for handling JWT-related errors.
Parameters:
  - secret:       A string representing the key for validating JWT tokens.

Returns:
  - middleware:   A Fiber middleware function for applying JWT protection to routes.
*/
func Protected(secret string) fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		Claims:       &presenter.JwtCustomClaims{},
		SigningKey:   []byte(secret),
		ErrorHandler: jwtError,
	})
}

/*
ValidateJWT is a Fiber middleware function that validates JWT tokens in incoming HTTP requests. It
attempts to validate the JWT token in the request context, and if successful, it sets the validated
user's ID in Fiber locals under the key "requester".
Parameters:
  - db:            GORM database instance for interacting with the database.
  - redisClient:   Redis client instance for additional token validation.
  - config:        JWT configuration settings.

Returns:
  - handler:    A Fiber middleware or handler function to serve HTTP requests.
*/
func ValidateJWT(db *gorm.DB, redisClient *redis.Client, _ jwtWare.Config) fiber.Handler {
	errMap := make(map[string]string)

	return func(c *fiber.Ctx) error {
		// Initialize a usecase instance with the provided database and Redis client
		uc := usecase.New(
			repository.New(db, accountRepo.New(db), courseRepo.New(db, accountRepo.New(db), courseGroupRepo.New(db))),
			accountRepo.New(db),
			courseRepo.New(db, accountRepo.New(db), courseGroupRepo.New(db)),
			redisClient)

		// Validate the JWT token in the request context
		user, err := uc.ValidateToken(c)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusForbidden).JSON(presenter.AuthErrorResponse(errMap))
		}

		access_key := "access_" + user.Username

		_, err = config.GetFromRedis(redisClient, access_key)
		if err != nil {
			errMap["error"] = "Invalid or expired access token."
			c.Status(http.StatusUnauthorized)
			return c.JSON(presenter.AuthErrorResponse(errMap))
		}
		// Set the validated user's ID in Fiber locals under the key "requesterID".
		c.Locals("requester", user.ID)

		// Continue with the next middleware or route handler
		return c.Next()
	}
}

/*
jwtError generates and returns a JSON response for handling JWT-related errors in a Fiber context.
Parameters:
  - c:   Fiber context representing the incoming HTTP request.
  - err: Error indicating the JWT-related issue.

Returns:
  - error: An error, if any occurred during the response generation.
*/
func jwtError(c *fiber.Ctx, err error) error {
	errMap := make(map[string]string)

	if err.Error() == "Missing or malformed JWT" {
		errMap["error"] = err.Error()
		return c.Status(http.StatusBadRequest).JSON(presenter.AuthErrorResponse(errMap))
	}

	errMap["error"] = err.Error()
	return c.Status(http.StatusUnauthorized).JSON(presenter.AuthErrorResponse(errMap))
}

/*
RolesAndPermissionMiddleware is a Fiber middleware function designed to enforce role-based access
control by checking the user's permissions before allowing access to a specific route or endpoint.
It takes a handler function as input, representing the subsequent middleware or endpoint to be executed.
The middleware first establishes a database connection, retrieves the user's ID from the Fiber context,
and utilizes an authentication handler to verify the user's permissions for the requested route.
If the user is a superadmin, they are granted access to all routes. Otherwise, the middleware checks
whether the user has the necessary permissions for the requested route, and access is granted accordingly.
If access is denied, an HTTP 403 Forbidden response is returned with appropriate error details.

Parameters:
  - handler: A Fiber handler function representing the subsequent middleware or endpoint to be executed
    if the user's permissions are validated by the middleware.

Returns:
  - fiber.Handler: A Fiber handler function that can be registered as middleware to enforce
    role-based access control for specific routes or endpoints.
*/
func RolesAndPermissionMiddleware(handler fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var err error

		// Get the global server configuration object
		server := config.AvyaasServer

		// Get the userID from c.Locals of Fiber context
		userID := c.Locals("requester").(uint)

		// Initialize the authHandler for accessing roles and per
		authHandler := core.GetAuth(server.DB)

		/* Get the url path of the requested page. Here, the second value returned by Split function
		is the url path stored in the database for the page.
		Example: c.Path() returns "/api/v1/role/list/". Hence, splitting by "/api/v1" will give "/role/list/",
		in the second index of permissionPath which is the url path stored in the database in our case.
		*/
		permissionPath := strings.Split(c.Path(), "/api/v3")

		// Check if the request user is administrator. If yes, has access to all url.
		isAdministrator, err := authHandler.CheckPermission(userID, 1)
		if isAdministrator && err == nil {
			return handler(c)
		}

		var valid bool
		// Check if the user has permission to access the url
		valid, _ = authHandler.CheckUserPermissionUrl(userID, permissionPath[1])

		if valid {
			return handler(c)
		}

		errMap["error"] = "forbidden, looks like you are trying to access wrong route"
		return c.Status(http.StatusForbidden).JSON(presenter.AuthErrorResponse(errMap))
	}
}

/*
RateLimiterMiddleware is a middleware function that limits the number of requests
a user can make within a specified time window.
*/
func RateLimiterMiddleware(maxRequests int, duration time.Duration) fiber.Handler {
	// Create a map to store the count of requests for each user
	var (
		mu       sync.Mutex
		counters = make(map[string]int)
	)

	return func(c *fiber.Ctx) error {
		// Get the user's IP address
		user := c.IP()

		mu.Lock()
		defer mu.Unlock()

		// Increment the request count for the user
		counters[user]++

		// Check if the user has exceeded the maximum allowed requests
		if counters[user] > maxRequests {
			errMap := make(map[string]string)
			errMap["error"] = "too many requests"

			fmt.Printf("errMap: %v\n", errMap)
			return c.Status(http.StatusBadRequest).JSON(errMap)
			// return errors.New("too many requests")
		}

		// Reset the request count after the specified duration
		time.AfterFunc(duration, func() {
			mu.Lock()
			defer mu.Unlock()

			counters[user] = 0
		})

		// Continue with the next handler
		return c.Next()
	}
}
