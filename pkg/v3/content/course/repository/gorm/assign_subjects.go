package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) AssignSubjectsToCourse(courseIds []uint, subjectIDs []uint) error {

	var courses []models.Course
	var subjects []models.Subject

	err := repo.db.Model(&models.Course{}).Where("id IN (?)", courseIds).Find(&courses).Error

	if err != nil {
		return err
	}

	err = repo.db.Model(&models.Subject{}).Where("id IN (?)", subjectIDs).Find(&subjects).Error

	if err != nil {
		return err
	}

	err = repo.db.Model(&courses).Association("Subjects").Append(subjects)

	if err != nil {
		return err

	}

	return nil
}
