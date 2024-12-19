package gorm

import (
	"avyaas/internal/domain/models"
	"errors"

	"gorm.io/gorm"
)

func (repo *Repository) GetUsersEnrolledInCourse(courseID uint) ([]models.User, error) {
	var enrolledUsers []models.User

	// Query the database to retrieve users enrolled in the course
	err := repo.db.Table("student_courses").
		Select("users.*").
		Joins("JOIN users ON student_courses.user_id = users.id").
		Where("student_courses.course_id = ?", courseID).
		Find(&enrolledUsers).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return enrolledUsers, nil
}
