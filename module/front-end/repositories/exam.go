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

func (r *Repository) CheckDuplicateQuestionRepo(ctx context.Context, question string) (bool, error) {
	var count int64
	if err := r.DB().Ctx().WithContext(ctx).
		Where("LOWER(question) = LOWER(?) AND deleted_at IS NULL", question).
		Model(&model.Exam{}).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Repository) CreateExamRepo(ctx context.Context, tx *gorm.DB, exam *model.Exam) error {
	tx = tx.WithContext(ctx)
	if err := tx.Create(exam).Error; err != nil {
		return err
	}
	return nil
}
