package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"

	"strings"
)

/*
RegisterStudent registers a new user by creating a user account with the provided information.

Parameters:
  - user: The user model containing registration information such as email, phone, password, etc.

Returns:
  - errMap:  A map of error messages, where keys represent the field names and values contain
    respective error details. If registration is successful, it returns nil.
*/
func (uCase *usecase) RegisterStudent(request presenter.StudentRegisterRequest) map[string]string {
	var err error
	errMap := make(map[string]string)
	var user models.User

	if _, err := uCase.accountRepo.GetUserByEmail(request.Identity); err == nil {
		errMap["email"] = fmt.Errorf("user with the email: '%s' already exists", request.Identity).Error()
	}

	if _, err := uCase.accountRepo.GetUserByPhone(request.Identity); err == nil {
		errMap["phone"] = fmt.Errorf("user with the phone: '%s' already exists", request.Identity).Error()
	}

	if len(errMap) != 0 {
		return errMap
	}

	if request.CourseID != 0 {
		if _, err := uCase.courseRepo.GetCourseByID(request.CourseID); err != nil {
			errMap["courseID"] = err.Error()
		}
	}

	if isValidEmail := utils.IsValidEmail(request.Identity); isValidEmail {
		splitEmail := strings.Split(request.Identity, "@")
		user.Username = strings.ToLower(splitEmail[0])

		user.Email = request.Identity
	} else if ok := utils.ContainsOnlyNumber(request.Identity); ok {
		if len(request.Identity) < 10 {
			errMap["identity"] = "phone number should be 10 digits"
			return errMap
		}
		user.Username = request.Identity

		user.Phone = request.Identity
	} else {
		errMap["error"] = "invalid email or phone"
		return errMap
	}

	// Derive username by splitting email or use phone number if email is empty
	// if user.Email != "" {
	// 	splitEmail := strings.Split(user.Email, "@")
	// 	user.Username = strings.ToLower(splitEmail[0])
	// } else {
	// 	user.Username = user.Phone
	// }

	// Check the strength of the provided password
	err = utils.CheckPasswordStrength(request.Password)
	if err != nil {
		errMap["password"] = err.Error()
		return errMap
	}

	// Hash the password for secure storage
	user.Password, err = utils.HashPassword(request.Password)
	if err != nil {
		errMap["password"] = err.Error()
		return errMap
	}

	// Assign a default student role ID = 4 to the user
	user.RoleID = 4
	user.CollegeName = request.CollegeName

	// Call the repository to save the user information
	if err = uCase.repository.RegisterStudent(user, request.Referral, request.CourseID); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return nil
}
