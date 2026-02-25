package services

import (
	"context"
	"errors"
	"strings"

	"github.com/tomatosAt/exam-08-golang/model"
	"github.com/tomatosAt/exam-08-golang/module/front-end/dto"
	"github.com/tomatosAt/exam-08-golang/module/front-end/mapper"
	"gorm.io/gorm"
)

func (s *Service) GetAllExamSVC(ctx context.Context) ([]dto.ExamResponse, error) {
	var examList []model.Exam
	if err := s.repo.GetAllExamRepo(ctx, nil, &examList); err != nil {
		return nil, err
	}
	response := mapper.MapListExamToResponse(examList)
	return response, nil
}

func (s *Service) CreateExamSVC(ctx context.Context, req dto.CreateExamRequest) (*dto.NewExamToResponse, error) {
	trimmedQuestion := strings.TrimSpace(req.Question)
	// เช็คคำถามซ้ำ
	isDuplicate, err := s.repo.CheckDuplicateQuestionRepo(ctx, trimmedQuestion)
	if err != nil {
		return nil, err
	}
	if isDuplicate {
		return nil, errors.New("question already exists")
	}
	// คำตอบถูกต้องต้องมีแค่ 1 ข้อเท่านั้น
	if err := s.CheckCorrectAnswer(req.Choices); err != nil {
		return nil, err
	}
	req.Question = trimmedQuestion
	exam := mapper.MapNewExam(req)
	// ใช้ tx commit กับ rollback เพราะ มีการ create 2 table พร้อมกัน ถ้าเกิด error ในขั้นตอนใดขั้นตอนหนึ่ง จะได้ rollback กลับทั้งหมด
	if err := s.repo.DB().Ctx().Transaction(func(tx *gorm.DB) error {
		return s.repo.CreateExamRepo(ctx, tx, &exam)
	}); err != nil {
		return nil, err
	}
	res := mapper.MapNewExamToResponse(&exam)
	return res, nil
}

func (s *Service) CheckCorrectAnswer(choices []dto.CreateChoiceRequest) error {
	correctCount := 0
	for _, c := range choices {
		if *c.IsCorrect {
			correctCount++
		}
	}
	if correctCount != 1 {
		return errors.New("คำตอบมีเพียง 1 ข้อเท่านั้น")
	}
	return nil
}

func (s *Service) DeleteExamSVC(ctx context.Context, id string) (*dto.DeleteExamResponse, error) {
	if err := s.repo.DB().Ctx().Transaction(func(tx *gorm.DB) error {
		return s.repo.DeleteExamRepo(ctx, tx, id)
	}); err != nil {
		return nil, err
	}
	res := mapper.MapDeleteExamToResponse(id)
	return res, nil
}
