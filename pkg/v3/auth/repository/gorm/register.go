package gorm

import (
	"avyaas/internal/core"
	"avyaas/internal/domain/models"
	"avyaas/utils"
	"fmt"
	"math/rand"

	"github.com/go-sql-driver/mysql"
)

/*
RegisterStudent persists a new user to the database and assigns a default role.

Parameters:
  - user: The user model containing registration information related to the user.

Returns:
  - err: An error if any step of the registration process fails; otherwise, returns nil.
*/
func (repo *Repository) RegisterStudent(user models.User, referral string, cID uint) error {
	// Initiates a database transaction for atomic operations
	transaction := repo.db.Begin()
	var existingUser models.User

	if err := repo.db.Where("username = ? ", user.Username).First(&existingUser).Error; err == nil {
		num := rand.Intn(100)
		user.Username = fmt.Sprintf("%s%d", user.Username, num)

		err = transaction.Create(&user).Error
		if err != nil {
			transaction.Rollback()
			return err
		}
	} else {
		err = transaction.Create(&user).Error
		if err != nil {
			transaction.Rollback()
			return err
		}
	}

	// Initialize the authority instance to access roles and permissions services
	auth := core.GetAuth(repo.db)

	// Assign the student role to the created user
	if err := auth.AssignRole(user.ID, 4); err != nil {
		transaction.Rollback()
		return err
	}

	// Generate a random OTP
	otp, err := utils.GenerateOTP()
	if err != nil {
		return err
	}

	var identity string
	if user.Email != "" {
		identity = user.Email
	} else {
		identity = user.Phone
	}

	// Save the OTP and user's phone number to the database
	err = repo.SaveUserOTP(identity, otp)
	if err != nil {
		return err
	}

	if isValidEmail := utils.IsValidEmail(identity); isValidEmail {
		// Send the OTP to the user in provided email
		if err = utils.SendOTPEmail(identity, otp); err != nil {
			return err
		}
	} else if ok := utils.ContainsOnlyNumber(identity); ok {
		// Send the OTP to the user in provided phone number
		if err = utils.SendOTPSMS(identity, otp); err != nil {
			return err
		}
	}

	student := &models.Student{
		Timestamp: models.Timestamp{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		ReferralCode: referral,
		// RegistrationNumber: regNumber,
	}

	if cID != 0 {
		regNumber, err := repo.GetRegistrationNumber(transaction, cID)
		if err != nil {
			transaction.Rollback()
			return fmt.Errorf("error generating registration number")
		}

		student.RegistrationNumber = regNumber
	}

	if referral != "" {
		teacher, err := repo.accountRepo.GetTeacherByReferralCode(referral)
		if err != nil {
			return fmt.Errorf("referral code: '%s' doesn't exists", referral)
		}

		err = transaction.Model(&models.Teacher{}).Where("id = ?", teacher.ID).Update("referral_count", teacher.ReferralCount+1).Error
		if err != nil {
			transaction.Rollback()
			return err
		}

		student.ReferredBy = teacher.ID
	}

	err = transaction.Create(&student).Error
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				newRegNumber, err := repo.GetRegistrationNumber(transaction, cID)
				if err != nil {
					transaction.Rollback()
					return fmt.Errorf("error generating registration number")
				}

				student.RegistrationNumber = newRegNumber
				newErr := transaction.Debug().Create(&student).Error
				if newErr != nil {
					transaction.Rollback()
					return newErr
				}

				transaction.Commit()
				return nil
			}

			transaction.Rollback()
			return err
		}
	}

	if cID != 0 {
		course, err := repo.courseRepo.GetCourseByID(cID)
		if err != nil {
			transaction.Rollback()
			return err
		}

		err = repo.courseRepo.EnrollInCourse(user.ID, cID)
		if err != nil {
			transaction.Rollback()
			return fmt.Errorf("cannot enroll in course: '%s'", course.Title)
		}
	}

	transaction.Commit()
	return err
}
