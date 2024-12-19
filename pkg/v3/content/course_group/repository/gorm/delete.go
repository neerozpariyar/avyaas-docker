package gorm

import (
	"avyaas/internal/domain/models"
)

/*
DeleteCourseGroup is a repository method responsible for deleting a course group with the specified
ID.

Parameters:
  - id: The ID of the course group to be deleted.

Returns:
  - err: An error, if any, encountered during the deletion operation.
*/
func (repo *Repository) DeleteCourseGroup(id uint) error {
	transaction := repo.db.Begin()
	courseGroup, err := repo.GetCourseGroupByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}

	if courseGroup.Thumbnail != "" {
		var cgFile models.File

		err = repo.db.Model(&models.File{}).Where("url = ?", courseGroup.Thumbnail).First(&cgFile).Error
		if err == nil {
			if err = transaction.Model(models.File{}).Where("id = ?", cgFile.ID).Update("is_active", false).Error; err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	err = transaction.Unscoped().Where("id = ?", id).Delete(&courseGroup).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
