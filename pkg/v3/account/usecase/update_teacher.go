package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"

	"strings"
)

/*
UpdateTeacher is a usecase function responsible for processing the update of teacher information.

Parameters:
  - uCase: A pointer to the use case struct, representing the business logic for user-related operations.
    It is used to access the repository for updating a teacher.
  - user: An instance of the TeacherCreateUpdateRequest struct containing the necessary information
    for updating a teacher.

Returns:
  - map[string]string: A map of error messages where keys represent specific fields or operations,
    and values contain corresponding error details. An empty map indicates a successful teacher update.
*/
func (uCase *usecase) UpdateTeacher(data presenter.TeacherCreateUpdateRequest) map[string]string {
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

	// if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
	// 	errMap["course"] = err.Error()
	// 	return errMap
	// }
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
	if err = uCase.repo.UpdateTeacher(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return nil
}
