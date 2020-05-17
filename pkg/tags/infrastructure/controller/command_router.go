package controller

import (
	"database/sql"
	"github.com/hernanhrm/memoriest/config"
	"github.com/hernanhrm/memoriest/pkg/tags/application"
	"github.com/hernanhrm/memoriest/pkg/tags/infrastructure/repository"
	"github.com/labstack/echo"
)

func CommandRoutes(app *echo.Echo, conf config.Model, db *sql.DB) {
	admin(app, conf, db)
}

func admin(app *echo.Echo, conf config.Model, db *sql.DB) {
	g := app.Group(adminPrefix)

	create(g, conf, db)
	update(g, conf, db)
	deleteR(g, conf, db)
}

func deleteR(g *echo.Group, conf config.Model, db *sql.DB) {
	repo := repository.GetDeleteByIDRepo(conf.PrimaryDBEngine, db)
	service := application.NewDeleteByIDService(repo)
	handle := newHandleDeleteByID(service)

	g.DELETE("/:id", handle.DeleteByID)
}

func update(g *echo.Group, conf config.Model, db *sql.DB) {
	repo := repository.GetUpdateRepo(conf.PrimaryDBEngine, db)
	service := application.NewServiceUpdateByID(repo)
	handle := newHandleUpdateByID(service)

	g.PUT("/:id", handle.Update)
}

func create(g *echo.Group, conf config.Model, db *sql.DB) {
	repo := repository.GetCreateRepo(conf.PrimaryDBEngine, db)
	service := application.NewCreateService(repo)
	handle := newHandleCreate(service)

	g.POST("", handle.Create)
}
