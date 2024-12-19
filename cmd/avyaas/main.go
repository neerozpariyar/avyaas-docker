/*
main is the entry point of the application, initializing the server, database, routes and other
required configurations.
*/

package main

import (
	"avyaas/cmd/avyaas/core"
	"avyaas/internal/config"

	"avyaas/pkg/v3/auth/handler/middleware"

	"log"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	jwtWare "github.com/gofiber/jwt/v2"
	"github.com/spf13/viper"
)

func main() {
	// Initiate viper configuration for reading config file
	config.ConfigureViper()

	// Initialize global server configuration parameters
	config.InitServerConfig()

	// Call global server config parameter variable
	server := config.AvyaasServer

	// Migrate database to server. Comment if migration not needed each time.
	// config.InitMigrations()

	app := server.App

	// NOTE: Uncomment the RateLimiterMiddleware function after testing later
	// Initiate request rate limiter middleware
	// app.Use(middleware.RateLimiterMiddleware(viper.GetInt("maxRequestLimit"), time.Second))

	// Logger middleware with configuration specified in logger.Config{} to log requests to the server
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${time} ${status} - ${method} ${path} ${time}\n",
	}))

	// Middleware for caching, compression, and logging in the order specified
	app.Use(
		logger.New(config.LoggerConfig()),
		compress.New(config.CompressResponseConfig()),
	)

	// CORS middleware to allow cross-origin requests
	server.App.Use(cors.New(config.CorsConfig()))

	// Configure authentication and authorization middleware
	authMW := middleware.Protected(viper.GetString("jwt.access_secret"))
	jwtMW := middleware.ValidateJWT(server.DB, server.RedisClient, jwtWare.Config{})

	/* authRoute is a Fiber.Router group created under the "/auth/v1/" path, to which authentication
	-related routes will be added */
	authRoute := app.Group("/auth/v3")

	/* protectedRoutes is a Fiber.Router group created under the "/api/v3/" path, to which other
	private routes will be added */
	protectedRoutes := server.App.Group("/api/v3", authMW, jwtMW)

	allRepos := core.InitRepositories(server)

	allUsecases := core.InitUsecases(server, *allRepos)

	core.InitRoutes(authRoute, protectedRoutes, allUsecases)

	// Start the Fiber web server on the specified port
	log.Fatal(app.Listen(":" + viper.GetString(`server.port`)))

}
