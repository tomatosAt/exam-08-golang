package frontend

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomatosAt/exam-08-golang/app"
	"github.com/tomatosAt/exam-08-golang/module/front-end/handler"
	"github.com/tomatosAt/exam-08-golang/module/front-end/repositories"
	s "github.com/tomatosAt/exam-08-golang/module/front-end/service"
)

func Create(app *app.Context) error {
	repo, err := repositories.New(app)
	if err != nil {
		return err
	}
	svc := s.New(repo)
	h := handler.NewHandler(svc)
	prefixPath := app.Config.App.PrefixPath + "-fe"
	g := app.Router.Group(prefixPath)
	addRouter(g, h)
	return nil
}

func addRouter(r fiber.Router, h *handler.Handler) {
	v1 := r.Group("/v1")
	exam := v1.Group("/exams")
	exam.Get("", h.GetAllExamHandler)
}
