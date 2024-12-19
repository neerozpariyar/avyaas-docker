package config

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

/*
FiberNewConf returns a new instance of Fiber's Config structure with the following pre-configured values:
- ServerHeader: "Avyaas"
- AppName: "Avyaas"
- DisableDefaultDate: true
- RequestMethods: GET, POST, HEAD, PATCH, DELETE, OPTIONS
- Concurrency: 256*1024*4
- JSONEncoder: json.Marshal
- JSONDecoder: json.Unmarshal

It takes no parameters and returns a fiber.Config structure.
*/
func NewFiberConfig() fiber.Config {
	conf := fiber.Config{
		ServerHeader:       "Avyaas",
		AppName:            "Avyaas",
		DisableDefaultDate: true,
		RequestMethods: []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPatch,
			fiber.MethodDelete,
			fiber.MethodOptions,
		},
		Concurrency: 256 * 1024 * 4,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   2 * 1024 * 1024 * 1024,
	}

	return conf
}

/*
CompressResponseConfig returns a compress.Config struct that is optimized for best compression.

Returns:

	A compress.Config struct with the following properties:
	- Level: compress.LevelBestCompression, which sets the compression level to the maximum
	  possible value for optimal compression.

Example usage:
  - config := CompressResponseConfig()
  - compressor := compress.New(config)
  - compressedData := compressor.Compress(data)
*/
func CompressResponseConfig() compress.Config {
	return compress.Config{
		Level: compress.LevelBestCompression,
	}
}

/*
LoggerConfig returns a logger configuration object with preconfigured options.

Returns:
  - logger.Config: The logger configuration object with preconfigured options

Example usage:
  - config := LoggerConfig()
*/
func LoggerConfig() logger.Config {
	return logger.Config{
		Format:     "[${ip}]:${port} ${time} ${status} - ${method} ${path} ${time}\n",
		TimeFormat: "02-12-2022",
		TimeZone:   "Asia/Kathmandu",
	}
}

/*
CorsConfig returns a cors.Config with predefined settings for Cross-Origin Resource Sharing (CORS).
It configures allowed headers, origins, credentials, and methods to facilitate secure cross-origin
requests. The returned cors.Config is intended for use in setting up CORS middleware in a web server.
*/
func CorsConfig() cors.Config {
	return cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}
}
