package http

import (
	"github.com/hernanhrm/memoriest/pkg/response"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func InitEcho() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	// e.Use(middleware.CORS())

	// CORS restricted
	// with GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: viper.GetStringSlice("allowed_domains"),
		AllowMethods: viper.GetStringSlice("allowed_methods"),
	}))

	e.HTTPErrorHandler = response.CustomHTTPErrorHandler

	return e
}
