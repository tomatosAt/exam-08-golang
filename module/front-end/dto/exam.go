package dto

type ExamResponse struct {
	ID       string           `json:"id"`
	Number   int              `json:"number"`
	Question string           `json:"question"`
	Choices  []ChoiceResponse `json:"choices"`
}
