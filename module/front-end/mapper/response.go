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
