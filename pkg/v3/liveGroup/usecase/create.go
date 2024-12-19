package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"

	"fmt"
)

func (uCase *usecase) CreateLiveGroup(data presenter.LiveGroupCreateUpdatePresenter) map[string]string {
	var err error
	errMap := make(map[string]string)

	if data.PackageTypeID != 0 {
		_, err := uCase.packageTypeRepo.GetPackageTypeByID(data.PackageTypeID)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}

		serviceIDs, err := uCase.packageTypeRepo.GetPackageTypeServices(data.PackageTypeID)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}

		if !utils.ContainsUint(serviceIDs, 4) {
			errMap["packageTypeID"] = "invalid package type ID"
			return errMap
		}
	}

	_, err = uCase.repo.GetLiveGroupByTitle(data.Title)
	if err == nil {
		errMap["title"] = fmt.Errorf("livegroup with title: '%s' already exists", data.Title).Error()
	}

	if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

	if errMap = uCase.repo.CreateLiveGroup(data); errMap != nil {
		return errMap
	}

	return errMap
}
