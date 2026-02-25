package model

import "github.com/google/uuid"

type Choice struct {
	Model
	ExamID     uuid.UUID `gorm:"type:char(36);not null"`
	ChoiceText string    `gorm:"size:255;not null"`
	IsCorrect  bool      `gorm:"default:false"`
}

func (Choice) TableName() string {
	return "choice"
}
