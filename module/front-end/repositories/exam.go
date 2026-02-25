package repositories

import (
	"context"

	"github.com/tomatosAt/exam-08-golang/model"
	"gorm.io/gorm"
)

func (r *Repository) GetAllExamRepo(ctx context.Context, tx *gorm.DB, examList *[]model.Exam) error {
	if tx == nil {
		tx = r.DB().Ctx().WithContext(ctx)
	}
	db := tx.
		Preload("Choices").
		Order("created_at asc")

	if err := db.Find(&examList).Error; err != nil {
		return err
	}
	return nil
}
