package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) GetVerifiedUsers() ([]models.User, error) {
	var users []models.User

	if err := repo.db.Where("verified = ?", true).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch verified users: %v", err)
	}

	return users, nil
}

func (repo *Repository) GetUnverifiedUsers() ([]models.User, error) {
	var users []models.User

	if err := repo.db.Where("verified = ?", false).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch verified users: %v", err)
	}

	return users, nil
}

func (repo *Repository) GetUsersByCourseID(courseID uint) ([]models.User, error) {
	var enrolledUsers []models.User

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
