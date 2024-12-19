package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"
)

/*
CreatePackage is a usecase method responsible for creating a new package.

Parameters:
  - data: A models.Package pointer instance representing the package to be created.

Returns:
  - errMap: A map[string]string containing error messages, if any, encountered during the process.
*/
func (uCase *usecase) CreatePackage(data presenter.PackageCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)
	// var course models.Course

	course, err := uCase.courseRepo.GetCourseByID(data.CourseID)
	if err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

	if _, err = uCase.packageTypeRepo.GetPackageTypeByID(data.PackageTypeID); err != nil {
		errMap["packageTypeID"] = err.Error()
		return errMap
	}

	if data.TestSeriesID != 0 {
		var testSeries *models.TestSeries
		if testSeries, err = uCase.testSeriesRepo.GetTestSeriesByID(data.TestSeriesID); err != nil {
			errMap["testSeriesID"] = err.Error()
			return errMap
		}

		if testSeries.CourseID != data.CourseID {
			errMap["testSeriesID"] = fmt.Errorf("test series: '%s' does not belong to the course: '%s'", testSeries.Title, course.CourseID).Error()
			return errMap
		}
	}

	if data.TestID != 0 {
		var test models.Test
		if test, err = uCase.testRepo.GetTestByID(data.TestID); err != nil {
			errMap["testID"] = err.Error()
			return errMap
		}

		if test.CourseID != data.CourseID {
			errMap["testID"] = fmt.Errorf("test: '%s' does not belong to the course: '%s'", test.Title, course.CourseID).Error()
			return errMap
		}
	}

	if data.LiveGroupID != 0 {
		var liveGroup models.LiveGroup
		liveGroup, err = uCase.liveGroupRepo.GetLiveGroupByID(data.LiveGroupID)
		if err != nil {
			errMap["liveGroupID"] = err.Error()
			return errMap
		}

		if liveGroup.CourseID != data.CourseID {
			errMap["liveGroupID"] = fmt.Errorf("live group: '%s' does not belong to the course: '%s'", liveGroup.Title, course.CourseID).Error()
			return errMap
		}
	}

	if data.LiveID != 0 {
		var live models.Live
		live, err = uCase.liveRepo.GetLiveByID(data.LiveID)
		if err != nil {
			errMap["liveID"] = err.Error()
			return errMap
		}

		if live.CourseID != data.CourseID {
			errMap["liveID"] = fmt.Errorf("live: '%s' does not belong to the course: '%s'", live.Topic, course.CourseID).Error()
			return errMap
		}
	}

	// for _, serviceID := range data.ServiceIDs {
	// 	_, err := uCase.serviceRepo.GetServiceByID(serviceID)
	// 	if err != nil {
	// 		errMap["serviceID"] = err.Error()
	// 	}
	// }

	// if len(errMap) != 0 {
	// 	return errMap
	// }

	// Call the repository to create the package
	if err = uCase.repo.CreatePackage(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
