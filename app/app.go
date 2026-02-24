package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/tomatosAt/exam-08-golang/config"
)

type Context struct {
	Config *config.Config
	Router *fiber.App
	log    *logrus.Entry
}

type App struct {
	*Context
}

func New(cfg *config.Config) *App {
	l := logrus.New()

	return &App{
		Context: &Context{
			Config: cfg,
			log:    l.WithField("package", "app"),
			Router: fiber.New(fiber.Config{
				AppName: cfg.App.Name,
			}),
		},
	}
}
