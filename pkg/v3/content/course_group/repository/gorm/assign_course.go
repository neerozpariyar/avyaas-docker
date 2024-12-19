package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) AssignCoursesToCourseGroup(courseGroupIds []uint, courseIds []uint) error {

	var courses []models.Course
	var courseGroups []models.CourseGroup

	err := repo.db.Model(&models.Course{}).Where("id IN (?)", courseIds).Find(&courses).Error

	if err != nil {
		return err
	}

	err = repo.db.Model(&models.CourseGroup{}).Where("id IN (?)", courseGroupIds).Find(&courseGroups).Error

	if err != nil {
		return err
	}

	err = repo.db.Model(&courseGroups).Association("Courses").Append(&courses)

	if err != nil {

		return err
	}

	return nil
}
