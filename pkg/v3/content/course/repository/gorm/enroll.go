package gorm

import (
	"avyaas/internal/domain/models"
	// "errors"
	// "gorm.io/gorm"
)

func (repo *Repository) EnrollInCourse(userID, courseID uint) error {
	var err error
	transaction := repo.db.Begin()

	// Check if user is already enrolled in the course
	// var enrolledCourse models.StudentCourse
	isPaid := false

	// if err = transaction.Where("user_id = ? AND course_id = ?", userID, courseID).First(&enrolledCourse).Error; err != nil {
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	err = transaction.Create(&models.StudentCourse{
		UserID:     userID,
		CourseID:   courseID,
		Paid:       &isPaid,
		ExpiryDate: nil,
	}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	// 	} else {
	// 		return err
	// 	}
	// }

	var course models.Course
	err = transaction.Where("id = ?", courseID).Preload("Subjects.Units.Chapters.Contents").First(&course).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	for _, subject := range course.Subjects {

		unitChapterContents, err := repo.GetRelationDataBySubject(subject.ID)
		if err != nil {
			return err
		}

		for _, unitChapterContent := range unitChapterContents {

			var content models.Content

			err := repo.db.Model(&models.Content{}).Where("id  = ?", unitChapterContent.ContentID).First(&content).Error

			if err != nil {
				transaction.Rollback()
				return err
			}
			isPaid := !(*content.IsPremium)

			studentContent := models.StudentContent{
				UserID:     userID,
				CourseID:   courseID,
				ContentID:  unitChapterContent.ContentID,
				Paid:       &isPaid,
				ExpiryDate: nil,
			}

			err = transaction.Create(&studentContent).Error

			if err != nil {
				transaction.Rollback()
				return err
			}

		}

	}

	transaction.Commit()
	return nil
}
