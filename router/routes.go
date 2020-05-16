package router

import (
	"database/sql"
	"github.com/gofiber/fiber"
	"github.com/hernanhrm/memoriest/config"
	"github.com/hernanhrm/memoriest/internal/user/infrastructure/http"
)

func Init(app *fiber.App, conf config.Model, db *sql.DB) {
	http.User(app, conf, db)
}
