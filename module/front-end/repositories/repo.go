package repositories

import (
	"github.com/sirupsen/logrus"
	"github.com/tomatosAt/exam-08-golang/app"
	"github.com/tomatosAt/exam-08-golang/config"
	"gorm.io/gorm"
)

const moduleName = "front-end"

type Repository struct {
	app    *app.Context
	log    *logrus.Entry
	dbMain *config.Client
}

func (r Repository) AppCfg() *config.Config {
	return r.app.Config
}

func (r Repository) Module() string {
	return moduleName
}

func (r Repository) Log() *logrus.Entry {
	return r.log.Dup()
}

func (r Repository) DB() *config.Client {
	return r.dbMain
}

func New(app *app.Context) (*Repository, error) {
	l := app.NewLogger().WithField("module", moduleName)
	dbMain := config.NewWithConfig(*app.Config, l.Logger)
	if err := dbMain.ConnectWithGormConfig(gorm.Config{}); err != nil {
		return nil, err
	}
	return &Repository{
		app:    app,
		log:    l,
		dbMain: dbMain,
	}, nil
}
