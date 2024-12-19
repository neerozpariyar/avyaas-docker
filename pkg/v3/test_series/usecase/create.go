package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"
)

func (uCase *usecase) CreateTestSeries(data presenter.TestSeriesCreateUpdateRequest) map[string]string {
	errMap := make(map[string]string)

	_, err := uCase.repo.GetTestSeriesByTitle(data.Title)
	if err == nil {
		errMap["title"] = fmt.Errorf("test series with title: '%s' already exists", data.Title).Error()
		return errMap
	}

	if data.IsPackage {
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

		if !utils.ContainsUint(serviceIDs, 2) {
			errMap["packageTypeID"] = "invalid package type ID"
			return errMap
		}

		if len(serviceIDs) > 1 {
			errMap["packageTypeID"] = "invalid package type ID"
			return errMap
		}
	}

	if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

	errMap = uCase.repo.CreateTestSeries(data)
	if errMap != nil {
		return errMap
	}

	return errMap

}
