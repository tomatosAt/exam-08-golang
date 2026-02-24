package app

import (
	"github.com/sirupsen/logrus"
)

func (ctx *Context) NewLogger() *logrus.Logger {
	l := logrus.New()
	return l
}
