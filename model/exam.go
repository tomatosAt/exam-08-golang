package model

import "gorm.io/gorm"

type Exam struct {
	Model
	Question  string `gorm:"type:varchar(255);not null"`
	DeletedAt gorm.DeletedAt
	Choices   []Choice `gorm:"foreignKey:ExamID;constraint:OnDelete:CASCADE"`
}

func (Exam) TableName() string {
	return "exam"
}
