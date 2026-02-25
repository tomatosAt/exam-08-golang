package services

import (
	"context"

	"github.com/tomatosAt/exam-08-golang/model"
	"github.com/tomatosAt/exam-08-golang/module/front-end/dto"
	"github.com/tomatosAt/exam-08-golang/module/front-end/mapper"
)

func (s *Service) GetAllExamSVC(ctx context.Context) ([]dto.ExamResponse, error) {
	var examList []model.Exam
	if err := s.repo.GetAllExamRepo(ctx, nil, &examList); err != nil {
		return nil, err
	}
	response := mapper.MapListExamToResponse(examList)
	return response, nil
}
