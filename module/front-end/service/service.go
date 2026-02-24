package services

import (
	"github.com/tomatosAt/exam-08-golang/module/front-end/ports"
	"github.com/tomatosAt/exam-08-golang/module/front-end/repositories"
)

type Service struct {
	repo ports.Repository
}

func New(repo *repositories.Repository) *Service {
	return &Service{repo: repo}
}
