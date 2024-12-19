package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetCourseByID(id uint) (models.Course, error) {
	var course models.Course

	// Retrieve the course from the database based on given id
	err := repo.db.Where("id = ?", id).First(&course).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Course{}, fmt.Errorf("course with course id: '%d' not found", id)
		}

		return models.Course{}, err
	}

	return course, nil
}

func (repo *Repository) GetCourseByCourseID(courseID string) (models.Course, error) {
	var course models.Course

	// Retrieve the course from the database based on given courseID
	err := repo.db.Where("course_id = ?", courseID).First(&course).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Course{}, err
		}

		return models.Course{}, err
	}

	return course, nil
}

func (repo *Repository) CheckStudentCourse(userID, courseID uint) (models.StudentCourse, error) {
	var course models.StudentCourse

	// Retrieve the course from the database based on given courseID
	err := repo.db.Where("user_id = ? And course_id = ?", userID, courseID).First(&course).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.StudentCourse{}, fmt.Errorf("not enrolled in the course")
		}

		return models.StudentCourse{}, err
	}

	return course, nil
}

func (repo *Repository) GetCourseGroupByCourseID(id uint) ([]models.CourseGroup, error) {
	var courseGroups []models.CourseGroup

	query := `SELECT course_groups.* FROM course_groups JOIN course_group_courses ON course_group_courses.course_group_id = course_groups.id WHERE course_group_courses.course_id  = ?`

	err := repo.db.Raw(query, id).Scan(&courseGroups).Error
	if err != nil {
		return nil, err
	}
	return courseGroups, nil
}

func (repo *Repository) GetRelationDataBySubject(subjectID uint) ([]models.UnitChapterContent, error) {
	var result []models.UnitChapterContent
	err := repo.db.Raw("SELECT * FROM unit_chapter_contents WHERE id IN (SELECT unit_chapter_content_id FROM subject_unit_chapter_contents WHERE subject_id  = ?)", subjectID).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *Repository) GetSubjectHeirarchy(id uint) ([]presenter.SubjectHeirarchyDetails, error) {

	var queryResults []presenter.SubjectHeirarchy
	var response []presenter.SubjectHeirarchyDetails
	var relationPrimaryKeys []uint

	err := repo.db.Select("unit_chapter_content_id").Table("subject_unit_chapter_contents").Where("subject_id  = ?", id).Find(&relationPrimaryKeys).Error

	if err != nil {
		return nil, err
	}

	query := `SELECT 
	ucc.id AS unit_chapter_content_id,
	u.id AS unit_id, u.title AS unit_title, u.description AS unit_description,
	ch.id AS chapter_id, ch.title AS chapter_title,
	c.id AS content_id, c.title AS content_title, c.is_premium AS content_is_premium, c.content_type AS content_type,
	FROM 
		unit_chapter_contents ucc
	JOIN 
		units u ON ucc.unit_id = u.id
	JOIN 
		chapters ch ON ucc.chapter_id = ch.id
	JOIN 
		contents c ON ucc.content_id = c.id
	WHERE 
		ucc.id IN (?);`

	err = repo.db.Raw(query, relationPrimaryKeys).Scan(&queryResults).Error

	if err != nil {
		return nil, err
	}

	for _, queryResult := range queryResults {
		singleDetail := presenter.SubjectHeirarchyDetails{
			ID: queryResult.UnitChapterContentID,
			Unit: presenter.FilterDetail{
				ID:    queryResult.UnitID,
				Title: queryResult.UnitTitle,
			},
			Chapter: presenter.FilterDetail{
				ID:    queryResult.ChapterID,
				Title: queryResult.ChapterTitle,
			},
			Content: presenter.FilterDetail{
				ID:    queryResult.ContentID,
				Title: queryResult.ContentTitle,
			},
		}

		response = append(response, singleDetail)
	}
	return response, nil
}
