package main

import (
	"github.com/hernanhrm/memoriest/config"
	"github.com/hernanhrm/memoriest/driver/database"
	"github.com/hernanhrm/memoriest/driver/http"
	"github.com/hernanhrm/memoriest/logger"
	"github.com/hernanhrm/memoriest/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	Run()
}

func Run() {
	conf := config.Model{}
	config.PopulateConfig("configuration", &conf, func() {
		log.Info("configuration settings reloaded")
	})

	logger.InitLogger(conf.LogFolder, conf.Env == "local")
	log.Info("Starting memoriest app")

	app := http.InitFiber(conf)

	err := database.GetPsqlConnection(conf)
	if err != nil {
		log.Fatalf("err connecting to psql database: %v", err)
	}
	db := database.GetConnection()

	router.Init(app, conf, db)

	err = app.Listen(8080)
	if err != nil {
		log.Fatalf("err when starting the server: %v", err)
	}
}
