package dto

type ChoiceResponse struct {
	ID         string `json:"id"`
	ChoiceText string `json:"choice_text"`
	IsCorrect  bool   `json:"is_correct"`
}

type CreateChoiceRequest struct {
	ChoiceText string `json:"choice_text" validate:"required"`
	IsCorrect  *bool  `json:"is_correct" validate:"required"`
}
