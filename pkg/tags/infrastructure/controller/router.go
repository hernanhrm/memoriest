package controller

import (
	"database/sql"
	"github.com/hernanhrm/memoriest/config"
	"github.com/labstack/echo"
)

const (
	adminPrefix  = "/api/v1/admin/tags"
	publicPrefix = "/api/v1/public/tags"
)

func Routes(app *echo.Echo, conf config.Model, db *sql.DB) {
	CommandRoutes(app, conf, db)
	QueryRoutes(app, conf, db)
}
