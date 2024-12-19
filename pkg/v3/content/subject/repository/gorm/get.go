package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/*
GetSubjectByID is a repository method responsible for retrieving a subject from the database
based on its unique identifier (ID).

Parameters:
  - id: A uint representing the unique identifier (ID) of the subject to be retrieved.

Returns:
  - subject: A models.Subject instance representing the retrieved subject.
  - error:    An error, if any, encountered during the database retrieval operation.
*/
func (repo *Repository) GetSubjectByID(id uint) (models.Subject, error) {
	var subject models.Subject

	// Retrieve the subject from the database based on given id
	err := repo.db.Where("id = ?", id).First(&subject).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Subject{}, fmt.Errorf("subject with subject id: '%d' not found", id)
		}

		return models.Subject{}, err
	}

	return subject, nil
}

/*
GetSubjectBySubjectID is a repository method responsible for retrieving a subject from the
database based on its subjectID.

Parameters:
  - subjectID: A string representing the unique identifier (subjectID) of the subject to be retrieved.

Returns:
  - subject: A models.Subject instance representing the retrieved subject.
  - error:    An error, if any, encountered during the database retrieval operation.
*/
func (repo *Repository) GetSubjectBySubjectID(subjectID string) (models.Subject, error) {
	var subject models.Subject

	// Retrieve the subject from the database based on given subjectID
	err := repo.db.Where("subject_id = ?", subjectID).First(&subject).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Subject{}, err
		}

		return models.Subject{}, err
	}

	return subject, nil
}

func (repo *Repository) GetCoursesBySubjectId(id uint) ([]models.Course, error) {
	var courses []models.Course

	query := `SELECT courses.* FROM courses JOIN course_subjects ON course_subjects.course_id = courses.id WHERE course_subjects.subject_id = ?`

	err := repo.db.Raw(query, id).Scan(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (repo *Repository) GetCourseIDsBySubjectID(id uint) ([]uint, error) {
	var courseIDs []uint

	err := repo.db.Raw("SELECT course_id FROM course_subjects WHERE subject_id  = ?", id).Scan(&courseIDs).Error

	if err != nil {
		return nil, err
	}
	return courseIDs, nil
}

func (repo *Repository) GetRelationsBySubjectID(id uint) ([]uint, error) {
	var relationIDs []uint

	err := repo.db.Select("unit_chapter_content_id").Table("subject_unit_chapter_contents").Where("subject_id  = ?", id).Find(&relationIDs).Error

	if err != nil {
		return nil, err
	}
	return relationIDs, nil
}

func (repo *Repository) CheckUnitInSubjectHeirarchy(subjectId uint) ([]uint, error) {
	var unitIDs []uint

	err := repo.db.Raw("SELECT unit_id FROM unit_chapter_contents WHERE id IN (SELECT unit_chapter_content_id FROM subject_unit_chapter_contents WHERE subject_id  = ?)", subjectId).Scan(&unitIDs).Error

	if err != nil {
		return nil, err
	}

	return unitIDs, err
}
