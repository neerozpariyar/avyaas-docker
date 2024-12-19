package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) CreateChapter(chapter models.Chapter) error {

	newChapter := &models.Chapter{
		Title: chapter.Title,
	}

	return repo.db.Create(&newChapter).Error
}
