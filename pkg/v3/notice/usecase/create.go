package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

func (uCase *usecase) CreateNotice(data presenter.NoticeCreateUpdatePresenter) map[string]string {
	var err error
	errMap := make(map[string]string)

	fmt.Printf("data.CourseID: %v\n", data.CourseID)
	if data.CourseID != 0 {
		_, err := uCase.courseRepo.GetCourseByID(uint(data.CourseID))
		if err != nil {
			errMap["courseID"] = err.Error()
			return errMap
		}
	}

	if err = uCase.repo.CreateNotice(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
