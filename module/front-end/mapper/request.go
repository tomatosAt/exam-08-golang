package mapper

import (
	"github.com/tomatosAt/exam-08-golang/model"
	"github.com/tomatosAt/exam-08-golang/module/front-end/dto"
)

func MapChoices(reqChoices []dto.CreateChoiceRequest) []model.Choice {
	choices := make([]model.Choice, 0)
	for _, c := range reqChoices {
		choices = append(choices, model.Choice{
			ChoiceText: c.ChoiceText,
			IsCorrect:  *c.IsCorrect,
		})
	}
	return choices
}

func MapNewExam(req dto.CreateExamRequest) model.Exam {
	return model.Exam{
		Question: req.Question,
		Choices:  MapChoices(req.Choices),
	}
}
