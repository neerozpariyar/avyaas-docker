package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"

	"errors"

	"gorm.io/gorm"
)

/*
GetUserByID retrieves a user by their unique ID from the database.

Parameters:
  - id: The unique identifier(ID) of the user to be retrieved.

Returns:
  - user: A pointer to an account.UserResponse containing the details of the retrieved user.
  - error: An error, if any occurred during the database query.
*/
func (repo *Repository) GetUserByID(id uint) (*presenter.UserResponse, error) {
	var user presenter.UserResponse

	if err := repo.db.Debug().Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		// If the user is not found, return a specific error message.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found for id: '%d'", id)
		}
		// Return other database-related errors as-is
		return nil, err
	}

	return &user, nil
}

/*
GetUserByEmail retrieves a user by their email from the database.
Parameters:
  - email: The email of the user to be retrieved.

Returns:
  - user: A pointer to an auth.UserResponse containing the details of the retrieved user.
  - error: An error, if any occurred during the database query.
*/
func (repo *Repository) GetUserByEmail(email string) (*presenter.UserResponse, error) {
	var user *presenter.UserResponse

	if err := repo.db.Debug().Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found for email: '%s'", email)
		}
		return nil, err
	}

	return user, nil
}

/*
GetUserByUsername retrieves a user by their unique username from the database.
Parameters:
  - username: The unique username of the user to be retrieved.

Returns:
  - user: A pointer to an presenter.UserResponse containing the details of the retrieved user.
  - error: An error, if any occurred during the database query.
*/
func (repo *Repository) GetUserByUsername(username string) (*presenter.UserResponse, error) {
	var user *presenter.UserResponse

	if err := repo.db.Debug().Table("users").Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found for username: '%s'", username)
		}
		return nil, err
	}

	return user, nil
}

/*
GetUserByPhone retrieves a user by their phone from the database.
Parameters:
  - phone: The phone number of the user to be retrieved.

Returns:
  - user: A pointer to an presenter.UserResponse containing the details of the retrieved user.
  - error: An error, if any occurred during the database query.
*/
func (repo *Repository) GetUserByPhone(phone string) (*presenter.UserResponse, error) {
	var user *presenter.UserResponse

	if err := repo.db.Debug().Table("users").Where("phone = ?", phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found for id: '%s'", phone)
		}
		return nil, err
	}

	return user, nil
}

func (repo *Repository) GetTeacherByID(id uint) (*models.Teacher, error) {
	var teacher *models.Teacher

	if err := repo.db.Debug().Where("id = ?", id).First(&teacher).Error; err != nil {
		// If the teacher is not found, return a specific error message.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("teacher not found for id: '%d'", id)
		}
		// Return other database-related errors as-is
		return nil, err
	}

	return teacher, nil
}

func (repo *Repository) GetTeacherByReferralCode(referral string) (*models.Teacher, error) {
	var teacher *models.Teacher

	if err := repo.db.Where("referral_code = ?", referral).First(&teacher).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("referral code not found:'%s'", referral)
		}
		return nil, err
	}
	return teacher, nil
}

func (repo *Repository) GetSubscriptionByUserID(userID uint) ([]models.Subscription, error) {
	var subscription []models.Subscription

	if err := repo.db.Where("user_id = ?", userID).Find(&subscription).Error; err != nil {
		return nil, err
	}
	return subscription, nil
}

func (repo *Repository) GetStudentsByTeacherReferral(id uint) ([]models.Student, error) {
	var students []models.Student

	if err := repo.db.Where("referred_by = ?", id).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (repo *Repository) GetStudentByID(id uint) (*models.Student, error) {
	var student models.Student

	if err := repo.db.Debug().Table("students").Where("id = ?", id).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("student not found for id: '%d'", id)
		}
		return nil, err
	}

	return &student, nil
}
