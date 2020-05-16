package http

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/limiter"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"github.com/hernanhrm/memoriest/config"
)

func InitFiber(c config.Model) *fiber.App {
	app := fiber.New(&fiber.Settings{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
	})

	// set the allowed origins and methods
	app.Use(cors.New(cors.Config{
		AllowOrigins: c.AllowedOrigins,
		AllowMethods: c.AllowedMethods,
	}))

	// limit the number of request
	app.Use(limiter.New())

	// logs request on stout
	app.Use(logger.New())

	// recover the app from unexpected errors
	app.Use(recover.New())

	// Helmet middleware provides protection against cross-site scripting (XSS) attack,
	// content type sniffing, clickjacking, insecure connection and other code injection attacks.
	app.Use(helmet.New())

	return app
}
