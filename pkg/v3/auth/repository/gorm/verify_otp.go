package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
	"log"

	"errors"
	"fmt"
	"time"

	authority "github.com/Ayata-Incorporation/roles_and_permission/cmd/roles_and_permissions"
	"gorm.io/gorm"
)

/*
VerifyUserOTP verifies the user-provided OTP (One-Time Password) for a given phone number.

Parameters:
  - otpRequest: The model containing the phone number and OTP provided by the user.

Returns:
  - verified: A boolean indicating whether the OTP verification is successful.
  - err: An error if any operation fails.
*/
func (repo *Repository) VerifyUserOTP(otpRequest models.UserOtp) (bool, error) {
	var userOtp models.UserOtp

	// Create a new instance of the authority.Authority struct
	auth := authority.New(authority.Options{DB: repo.db})

	// Queries the database for the stored OTP record associated with the provided phone number
	err := repo.db.Model(&models.UserOtp{}).Where("identity = ?", otpRequest.Identity).First(&userOtp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf("otp not found for : '%s'", otpRequest.Identity)
		}
		return false, err
	}

	// Checks if the OTP matches the one provided by the user and is not expired
	if otpRequest.OTP != userOtp.OTP || userOtp.ExpiresAt.Unix() < time.Now().Unix() {
		return false, errors.New("expired or wrong otp")
	}

	var user *models.User
	// Updates the user record in the database, marking it as verified
	// err = repo.db.Model(&models.User{}).Where("phone = ?", otpRequest.Identity).Update("verified", true).Scan(&user).Error
	// if err != nil {
	// 	return false, err
	// }

	var baseQuery *gorm.DB

	if isValidEmail := utils.IsValidEmail(otpRequest.Identity); isValidEmail {
		baseQuery = repo.db.Model(&models.User{}).Where("email = ?", otpRequest.Identity).Scan(&user)
		// err = repo.db.Model(&models.User{}).Where("email = ?", otpRequest.Identity).Update("verified", true).Scan(&user).Error
		// if err != nil {
		// 	return false, err
		// }
	} else if ok := utils.ContainsOnlyNumber(otpRequest.Identity); ok {
		baseQuery = repo.db.Model(&models.User{}).Where("phone = ?", otpRequest.Identity).Scan(&user)
		// err = repo.db.Model(&models.User{}).Where("phone = ?", otpRequest.Identity).Update("verified", true).Scan(&user).Error
		// if err != nil {
		// 	return false, err
		// }
	}

	err = baseQuery.Update("verified", true).Error
	if err != nil {
		return false, err
	}

	if !user.Verified && user.RoleID == 4 {
		// Assign permission with id "1" which gives all permissionto the created user
		if errList := auth.AssignUserPermissions(
			user.ID, []uint{3, 8, 10, 14, 15, 18, 19, 21, 22, 26, 31, 36, 37, 42, 43, 45, 49, 53, 57, 58,
				63, 64, 65, 67, 68, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92,
				93, 95, 99, 100, 104, 108, 112, 114, 118, 121, 127, 130, 131, 132, 134, 135, 137, 139}); len(errList) != 0 {
			log.Println(errList)
		}
	}

	return true, nil
}

// []uint{3, 6, 8, 10, 14, 15, 21, 22, 26, 31, 36, 37, 45, 57, 58, 64, 65, 67, 68,
//                 73, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93,
//                 104, 134, 135, 136, 137, 140, 141, 142, 143}

// 				10:  {"List course groups", "View a list of all course groups"},
// 				,		14:  {"List courses", "View a list of all courses"},
// 				,		15:  {"View course details", "View detailed information about a course including all the contents inside"},
// 				,		90:  {"Create feedback", "Create a new feedback"},
// 				,		91:  {"List feedbacks", "View a list of all feedbacks"},
// 				,		92:  {"Update feedback", "Update feedback information"},
// 				,		93:  {"Delete feedback", "Delete a feedback"},
// 				,		134: {"Create bookmark", "Create a new bookmark"},
// 				,		135: {"List bookmark", "View a list of all bookmarks"},
// 				,		136: {"View bookmark details", "View detailed information about a bookmark"},
// 				,		137: {"Delete bookmark", "Delete a bookmark"},
// 				,

// 				76:  {"Create discussion", "Create a new discussion"},
