package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"

	"strings"
)

/*
CreateTeacher is a usecase function responsible for processing the creation of a new teacher.

Parameters:
  - uCase: A pointer to the use case struct, representing the business logic for user-related operations.
    It is used to access the repository for creating a new teacher.
  - user: An instance of the TeacherCreateUpdateRequest struct containing the necessary information
    for creating a new teacher.

Returns:
  - map[string]string: A map of error messages where keys represent specific fields or operations,
    and values contain corresponding error details. An empty map indicates a successful teacher creation.
*/
func (uCase *usecase) CreateTeacher(data presenter.TeacherCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	if _, err := uCase.repo.GetUserByEmail(data.Email); err == nil {
		errMap["email"] = fmt.Errorf("teacher with email: '%s' already exists", data.Email).Error()
		return errMap
	}

	if _, err := uCase.repo.GetUserByPhone(data.Phone); err == nil {
		errMap["phone"] = fmt.Errorf("teacher with phone: '%s' already exists", data.Phone).Error()
		return errMap
	}

	for _, subjectID := range data.SubjectIDs {
		_, err := uCase.subjectRepo.GetSubjectByID(subjectID)
		if err != nil {
			errMap["error"] = err.Error()
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
	if err = uCase.repo.CreateTeacher(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return nil
}
