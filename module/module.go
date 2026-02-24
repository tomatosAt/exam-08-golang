package module

import (
	"github.com/tomatosAt/exam-08-golang/app"
	frontend "github.com/tomatosAt/exam-08-golang/module/front-end"
)

func Create(a *app.App) error {
	l := a.NewLogger().WithField("module", "global")
	if err := frontend.Create(a.Context); err != nil {
		l.Errorln("[x] Create FrontEndAPI module error -:", err)
		return err
	}
	return nil
}
