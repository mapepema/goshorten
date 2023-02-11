package secure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CORS adds Corss-Origin Resource Sharing support
func CORS() fiber.Handler {
	return cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "*",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length",
		MaxAge:           86400,
	})
}
