package model

import "gorm.io/gorm"

type Choice struct {
	gorm.Model
	ExamID     uint   `gorm:"not null"`
	ChoiceText string `gorm:"size:255;not null"`
	IsCorrect  bool   `gorm:"default:false"`
}

func (Choice) TableName() string {
	return "choice"
}
