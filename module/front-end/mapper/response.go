package mapper

import (
	"github.com/tomatosAt/exam-08-golang/model"
	"github.com/tomatosAt/exam-08-golang/module/front-end/dto"
)

func MapListExamToResponse(e []model.Exam) []dto.ExamResponse {
	var response []dto.ExamResponse

	for i, exam := range e {

		var choices []dto.ChoiceResponse
		for _, c := range exam.Choices {
			choices = append(choices, dto.ChoiceResponse{
				ID:         c.ID.String(),
				ChoiceText: c.ChoiceText,
				IsCorrect:  c.IsCorrect,
			})
		}

		response = append(response, dto.ExamResponse{
			ID:       exam.ID.String(),
			Number:   i + 1,
			Question: exam.Question,
			Choices:  choices,
		})
	}

	return response
}

func MapNewExamToResponse(e *model.Exam) *dto.NewExamToResponse {

	choices := make([]dto.ChoiceResponse, 0)

	for _, c := range e.Choices {
		choices = append(choices, dto.ChoiceResponse{
			ID:         c.ID.String(),
			ChoiceText: c.ChoiceText,
			IsCorrect:  c.IsCorrect,
		})
	}

	return &dto.NewExamToResponse{
		ID:       e.ID.String(),
		Question: e.Question,
		Choices:  choices,
	}
}

func MapDeleteExamToResponse(id string) *dto.DeleteExamResponse {
	return &dto.DeleteExamResponse{
		ID:      id,
		Message: "Delete success",
	}
}
