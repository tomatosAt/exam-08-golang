package dto

type ExamResponse struct {
	ID       string           `json:"id"`
	Number   int              `json:"number"`
	Question string           `json:"question"`
	Choices  []ChoiceResponse `json:"choices"`
}

type CreateExamRequest struct {
	Question string                `json:"question" validate:"required"`
	Choices  []CreateChoiceRequest `json:"choices" validate:"required,min=1,dive"`
}

type NewExamToResponse struct {
	ID       string           `json:"id"`
	Question string           `json:"question"`
	Choices  []ChoiceResponse `json:"choices"`
}
