package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
	"math"

	"github.com/spf13/viper"
)

func (repo *Repository) ListEnrolledCourse(userID uint, page int, search string, pageSize int) ([]models.Course, float64, error) {
	var courses []models.Course

	err := repo.db.Debug().Model(&models.Course{}).Where("id IN (?)",
		repo.db.Select("course_id").Model(&models.StudentCourse{}).Where("user_id = ?", userID)).Scopes(utils.Paginate(page, pageSize)).Order("id").Find(&courses).Error

	totalPage := math.Ceil(float64(len(courses)) / float64(viper.GetInt("pagination.page_size")))

	return courses, totalPage, err
}
