package main

import (
	"fmt"
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

	// app := http.InitFiber(conf)
	app := http.InitEcho()
	err := database.GetPsqlConnection(conf)
	if err != nil {
		log.Fatalf("err connecting to psql database: %v", err)
	}
	db := database.GetConnection()

	router.Init(app, conf, db)

	port := fmt.Sprintf(":%d", conf.PortHTTP)

	log.Fatal(app.Start(port))
}
