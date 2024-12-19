package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"fmt"
)

func (uCase *usecase) UpdateTestSeries(data presenter.TestSeriesCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing test series with the provided test series's ID
	ts, err := uCase.repo.GetTestSeriesByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Check if a testSeries with the given title already exists
	tsByTitle, err := uCase.repo.GetTestSeriesByTitle(data.Title)
	if err == nil {
		// Check if the title is the same as of the requested test series
		if ts.Title != tsByTitle.Title {
			errMap["title"] = fmt.Errorf("test series with title: '%s' already exists", tsByTitle.Title).Error()
			return errMap
		}
	}

	if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

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

		if !utils.ContainsUint(serviceIDs, 2) {
			errMap["packageTypeID"] = "invalid package type ID"
			return errMap
		}

		if len(serviceIDs) > 1 {
			errMap["packageTypeID"] = "invalid package type ID"
			return errMap
		}
	}

	// Delegate the update of test series
	if errMap = uCase.repo.UpdateTestSeries(data); errMap != nil {
		return errMap
	}

	return errMap
}
