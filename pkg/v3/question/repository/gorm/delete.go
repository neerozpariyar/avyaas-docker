package gorm

import (
	"avyaas/internal/domain/models"
)

/*
DeleteQuestion is a repository method responsible for deleting a question with the specified ID.

Parameters:
  - id: The ID of the question to be deleted.

Returns:
  - err: An error, if any, encountered during the deletion operation.
*/
func (repo *Repository) DeleteQuestion(id uint) error {
	transaction := repo.db.Begin()

	options, err := repo.GetOptionsByQuestionID(id)
	if err != nil {
		return err
	}

	for _, option := range options {
		if option.Image != "" {
			var imageFile models.File

			err = repo.db.Model(&models.File{}).Where("url = ?", option.Image).First(&imageFile).Error
			if err == nil {
				if err = transaction.Model(models.File{}).Where("id = ?", imageFile.ID).Update("is_active", false).Error; err != nil {
					transaction.Rollback()
					return err
				}
			}
		}

		if option.Audio != "" {
			var audioFile models.File

			err = repo.db.Model(&models.File{}).Where("url = ?", option.Audio).First(&audioFile).Error
			if err == nil {
				if err = transaction.Model(models.File{}).Where("id = ?", audioFile.ID).Update("is_active", false).Error; err != nil {
					transaction.Rollback()
					return err
				}
			}
		}
	}

	err = transaction.Model(&models.Option{}).Where("question_id = ?", id).Delete(&models.Option{}).Error

	if err != nil {
		return err
	}
	// option, _ := repo.GetOptionByQuestionID(id)
	// if option != nil {
	// 	// Option A
	// 	if option.Image != nil {
	// 		var imageFile models.File

	// 		err = repo.db.Model(&models.File{}).Where("url = ?", option.ImageA).First(&imageFile).Error
	// 		if err == nil {
	// 			if err = transaction.Model(models.File{}).Where("id = ?", imageFile.ID).Update("is_active", false).Error; err != nil {
	// 				transaction.Rollback()
	// 				return err
	// 			}
	// 		}
	// 	}

	// 	if option.Audio != nil {
	// 		var audioFile models.File

	// 		err = repo.db.Model(&models.File{}).Where("url = ?", option.AudioA).First(&audioFile).Error
	// 		if err == nil {
	// 			if err = transaction.Model(models.File{}).Where("id = ?", audioFile.ID).Update("is_active", false).Error; err != nil {
	// 				transaction.Rollback()
	// 				return err
	// 			}
	// 		}
	// 	}

	// 	// Option B
	// 	if option.ImageB != "" {
	// 		var imageFile models.File

	// 		err = repo.db.Model(&models.File{}).Where("url = ?", option.ImageB).First(&imageFile).Error
	// 		if err == nil {
	// 			if err = transaction.Model(models.File{}).Where("id = ?", imageFile.ID).Update("is_active", false).Error; err != nil {
	// 				transaction.Rollback()
	// 				return err
	// 			}
	// 		}
	// 	}

	// 	if option.AudioB != "" {
	// 		var audioFile models.File

	// 		err = repo.db.Model(&models.File{}).Where("url = ?", option.AudioB).First(&audioFile).Error
	// 		if err == nil {
	// 			if err = transaction.Model(models.File{}).Where("id = ?", audioFile.ID).Update("is_active", false).Error; err != nil {
	// 				transaction.Rollback()
	// 				return err
	// 			}
	// 		}
	// 	}

	// 	// Option C
	// 	if option.ImageC != "" {
	// 		var imageFile models.File

	// 		err = repo.db.Model(&models.File{}).Where("url = ?", option.ImageC).First(&imageFile).Error
	// 		if err == nil {
	// 			if err = transaction.Model(models.File{}).Where("id = ?", imageFile.ID).Update("is_active", false).Error; err != nil {
	// 				transaction.Rollback()
	// 				return err
	// 			}
	// 		}
	// 	}

	// 	if option.AudioC != "" {
	// 		var audioFile models.File

	// 		err = repo.db.Model(&models.File{}).Where("url = ?", option.AudioC).First(&audioFile).Error
	// 		if err == nil {
	// 			if err = transaction.Model(models.File{}).Where("id = ?", audioFile.ID).Update("is_active", false).Error; err != nil {
	// 				transaction.Rollback()
	// 				return err
	// 			}
	// 		}
	// 	}

	// 	// Option D
	// 	if option.ImageD != "" {
	// 		var imageFile models.File

	// 		err = repo.db.Model(&models.File{}).Where("url = ?", option.ImageB).First(&imageFile).Error
	// 		if err == nil {
	// 			if err = transaction.Model(models.File{}).Where("id = ?", imageFile.ID).Update("is_active", false).Error; err != nil {
	// 				transaction.Rollback()
	// 				return err
	// 			}
	// 		}
	// 	}

	// 	if option.AudioD != "" {
	// 		var audioFile models.File

	// 		err = repo.db.Model(&models.File{}).Where("url = ?", option.AudioD).First(&audioFile).Error
	// 		if err == nil {
	// 			if err = transaction.Model(models.File{}).Where("id = ?", audioFile.ID).Update("is_active", false).Error; err != nil {
	// 				transaction.Rollback()
	// 				return err
	// 			}
	// 		}
	// 	}
	// }

	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.Question{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
