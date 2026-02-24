package model

import "gorm.io/gorm"

type Exam struct {
	gorm.Model
	Number   int      `gorm:"not null;index"`
	Question string   `gorm:"type:varchar(255);not null"`
	Choices  []Choice `gorm:"foreignKey:ExamID;constraint:OnDelete:CASCADE"`
}

func (Exam) TableName() string {
	return "exam"
}
