package usecase

import (
	"avyaas/internal/domain/models"
)

func (uCase *usecase) CreateNotification(data models.Notification) map[string]string {
	var err error
	errMap := make(map[string]string)

	// if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
	// 	errMap["courseID"] = err.Error()
	// 	return errMap
	// }
	// if _, err := uCase.accountRepo.GetUserByID(data.UserID); err != nil {
	// 	errMap["userID"] = err.Error()
	// 	return errMap
	// }

	if data.Recipient == "course" {
		if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
			errMap["courseID"] = err.Error()
			return errMap
		}
	}

	if err = uCase.repo.CreateNotification(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
