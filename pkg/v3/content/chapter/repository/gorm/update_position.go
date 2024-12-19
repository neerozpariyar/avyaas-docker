package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (repo *Repository) UpdateChapterPosition(data presenter.UpdateChapterPositionRequest) map[string]string {
	errMap := make(map[string]string)
	transaction := repo.db.Begin()

	for idx, chapterID := range data.ChapterIDs {
		_, err := repo.GetChapterByID(chapterID)
		if err != nil {
			errMap["chapterID"] = err.Error()
		}

		err = transaction.Model(&models.Chapter{}).Where("id = ?", chapterID).Update("position", idx+1).Error
		if err != nil {
			errMap["chapterID"] = err.Error()
		}
	}

	if len(errMap) != 0 {
		transaction.Rollback()
		return errMap
	}

	transaction.Commit()
	return errMap
}
