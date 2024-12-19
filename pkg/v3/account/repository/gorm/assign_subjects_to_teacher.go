package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
)

func (repo *Repository) AssignSubjectsToTeacher(userID uint, subjectIDs []uint) error {
	var oldSubjectIDs []uint
	transaction := repo.db.Begin()

	err := repo.db.Model(&models.TeacherSubject{}).Select("subject_id").Where("user_id = ?", userID).Scan(&oldSubjectIDs).Error
	if err != nil {
		return err
	}
	addSubjects, delSubjects := utils.CompareDifferences(oldSubjectIDs, subjectIDs)
	for _, subjectID := range addSubjects {
		var courseID uint
		err := repo.db.Debug().Model(&models.Subject{}).Select("course_id").Where("id= ?", subjectID).Scan(&courseID).Error
		if err != nil {
			return err
		}
		teacherSubject := models.TeacherSubject{
			UserID:    userID,
			SubjectID: subjectID,
			CourseID:  courseID,
		}

		err = transaction.Create(&teacherSubject).Error
		if err != nil {
			transaction.Rollback()
			return err
		}

	}
	if len(delSubjects) > 0 {
		for _, subjectID := range delSubjects {
			err = transaction.Unscoped().Where("user_id = ? AND subject_id = ?", userID, subjectID).Delete(&models.TeacherSubject{}).Error

			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	err = transaction.Commit().Error

	if err != nil {
		transaction.Rollback()
		return err
	}

	return nil
}
