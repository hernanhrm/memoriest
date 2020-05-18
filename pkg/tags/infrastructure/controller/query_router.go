package controller

import (
	"database/sql"
	"github.com/hernanhrm/memoriest/config"
	"github.com/hernanhrm/memoriest/pkg/tags/application"
	"github.com/hernanhrm/memoriest/pkg/tags/infrastructure/repository"
	"github.com/labstack/echo"
)

func QueryRoutes(app *echo.Echo, conf config.Model, db *sql.DB) {
	g := app.Group(publicPrefix)

	public(g, conf, db)
}

func public(g *echo.Group, conf config.Model, db *sql.DB) {
	getAllRoute(g, conf, db)
	getByIDRoute(g, conf, db)
}

func getByIDRoute(g *echo.Group, conf config.Model, db *sql.DB) {
	repo := repository.GetGetByID(conf.PrimaryDBEngine, db)
	service := application.NewServiceGetByID(repo)
	handle := newHandleGetByID(service)

	g.GET("/:id", handle.GetByID)
}

func getAllRoute(g *echo.Group, conf config.Model, db *sql.DB) {
	repo := repository.GetGetAll(conf.PrimaryDBEngine, db)
	service := application.NewServiceGetAll(repo)
	handle := NewHandleGetAll(service)

	g.GET("", handle.GetAll)
}
