package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) DeleteCourse(id uint) error {
	transaction := repo.db.Begin()

	course, err := repo.GetCourseByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}

	if course.Thumbnail != "" {
		var cFile models.File

		err = repo.db.Model(&models.File{}).Where("url = ?", course.Thumbnail).First(&cFile).Error
		if err == nil {
			if err = transaction.Model(models.File{}).Where("id = ?", cFile.ID).Update("is_active", false).Error; err != nil {
				transaction.Rollback()
				return err
			}
		}
	}
	// Perform a hard delete of the course with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&course).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
