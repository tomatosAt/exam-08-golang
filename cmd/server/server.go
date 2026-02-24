package main

import (
	"log"
	"os"

	"github.com/tomatosAt/exam-08-golang/app"
	"github.com/tomatosAt/exam-08-golang/config"
	"github.com/tomatosAt/exam-08-golang/module"
)

func startServer() {
	cfg := config.LoadConfig("config.yml")
	a := app.New(cfg)
	// router
	if err := module.Create(a); err != nil {
		log.Println("[x] Start create module failed -:", err)
		os.Exit(1)
	}

	if err := a.Router.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
