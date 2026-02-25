package dto

type ChoiceResponse struct {
	ID         string `json:"id"`
	ChoiceText string `json:"choice_text"`
	IsCorrect  bool   `json:"is_correct"`
}
