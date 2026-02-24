package ports

import (
	"github.com/sirupsen/logrus"
	"github.com/tomatosAt/exam-08-golang/config"
)

type Repository interface {
	AppCfg() *config.Config
	Module() string
	Log() *logrus.Entry
	DB() *config.Client
}

type Service interface {
}
