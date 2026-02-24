package handler

import (
	"github.com/tomatosAt/exam-08-golang/module/front-end/ports"
	s "github.com/tomatosAt/exam-08-golang/module/front-end/service"
)

type Handler struct {
	svc ports.Service
}

func NewHandler(svc *s.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}
