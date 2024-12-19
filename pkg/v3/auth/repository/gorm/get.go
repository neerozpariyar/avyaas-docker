package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"

	"errors"

	"gorm.io/gorm"
)

/*
GetRoleByID retrieves a user role by its unique ID from the database.
Parameters:
  - roleID: The unique identifier of the user role to be retrieved.

Returns:
  - role: A pointer to an auth.UserRoleResponse presenter containing the details of the retrieved user role.
  - error: An error, if any occurred during the database query.
*/
func (r *Repository) GetRoleByID(roleID int) (*presenter.UserRoleResponse, error) {
	var role *presenter.UserRoleResponse

	if err := r.db.Table("roles").Where("id = ?", roleID).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid role id")
		}
		return nil, err
	}

	return role, nil
}

func (repo *Repository) GetRegistrationNumber(transaction *gorm.DB, cID uint) (string, error) {
	var registrationNumber *models.RegistrationCount
	var newRegNum uint
	var course models.Course

	if err := repo.db.Where("id = ?", cID).First(&course).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("course registration not found for course: '%s'", course.CourseID)
		}
		return "", err
	}

	if err := repo.db.Where("course_id = ?", cID).First(&registrationNumber).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("course ID not found")
		}
		return "", err
	}

	fmt.Printf("registrationNumber.RegistrationCount: %v\n", registrationNumber.RegistrationCount)
	if err := transaction.Model(&models.RegistrationCount{}).Select("registration_count").Where("course_id = ?", cID).Update("registration_count", registrationNumber.RegistrationCount+1).Scan(&newRegNum).Error; err != nil {
		return "", err
	}

	regNum := course.CourseID + "-" + utils.UintToString(newRegNum)

	fmt.Printf("registrationNumber: %v\n", newRegNum)

	return regNum, nil
}
