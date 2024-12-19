package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"
	"strings"
)

// update profile by self only

func (uCase *usecase) UpdateStudent(data presenter.StudentCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	userByID, err := uCase.repo.GetUserByID(data.ID)
	if err != nil {
		errMap["email"] = err.Error()
		return errMap
	}

	userByEmail, err := uCase.repo.GetUserByEmail(data.Email)
	if err == nil {
		if userByID.ID != userByEmail.ID {
			errMap["email"] = fmt.Errorf("user with email: '%s' already exists", userByEmail.Email).Error()
			return errMap
		}
	}

	userByPhone, err := uCase.repo.GetUserByPhone(data.Phone)
	if err == nil {
		if userByID.ID != userByPhone.ID {
			errMap["phone"] = fmt.Errorf("user with phone: '%s' already exists", userByPhone.Phone).Error()
			return errMap
		}
	}

	// Derive username by splitting email or use phone number if email is empty
	if data.Email != "" {
		splitEmail := strings.Split(data.Email, "@")
		data.Username = strings.ToLower(splitEmail[0])
	} else {
		data.Username = data.Phone
	}

	gender := []string{"male", "female", "other"}
	if !utils.Contains(gender, strings.ToLower(string(data.Gender))) {
		errMap["gender"] = "invalid gender"
		return errMap
	}

	data.Gender = strings.ToUpper(string(data.Gender))

	// Call the repository to save the user information
	if err = uCase.repo.UpdateStudent(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return nil
}
