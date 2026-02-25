package ports

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/tomatosAt/exam-08-golang/config"
	"github.com/tomatosAt/exam-08-golang/model"
	"github.com/tomatosAt/exam-08-golang/module/front-end/dto"
	"gorm.io/gorm"
)

type Repository interface {
	AppCfg() *config.Config
	Module() string
	Log() *logrus.Entry
	DB() *config.Client
	GetAllExamRepo(ctx context.Context, tx *gorm.DB, examList *[]model.Exam) error
	CreateExamRepo(ctx context.Context, tx *gorm.DB, exam *model.Exam) error
	CheckDuplicateQuestionRepo(ctx context.Context, question string) (bool, error)
}

type Service interface {
	GetAllExamSVC(ctx context.Context) ([]dto.ExamResponse, error)
	CreateExamSVC(ctx context.Context, req dto.CreateExamRequest) (*dto.NewExamToResponse, error)
}
