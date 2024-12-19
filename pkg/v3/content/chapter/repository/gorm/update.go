package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) UpdateChapter(chapter models.Chapter) error {
	return repo.db.Where("id = ?", chapter.ID).Updates(&chapter).Error
}
