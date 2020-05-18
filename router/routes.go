package router

import (
	"database/sql"
	"github.com/hernanhrm/memoriest/config"
	tags "github.com/hernanhrm/memoriest/pkg/tags/infrastructure/controller"
	"github.com/labstack/echo"
)

func Init(app *echo.Echo, conf config.Model, db *sql.DB) {
	tags.Routes(app, conf, db)
}
