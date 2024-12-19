package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (u *usecase) ListPackage(request *presenter.PackageListRequest) ([]presenter.PackageListResponse, int, error) {
	packages, totalPage, err := u.repo.ListPackage(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allPackages []presenter.PackageListResponse
	for _, pkg := range packages {
		eachPackage := presenter.PackageListResponse{
			ID:          pkg.ID,
			Title:       pkg.Title,
			Description: pkg.Description,
			Price:       pkg.Price,
			Period:      pkg.Period,
		}

		course, err := u.courseRepo.GetCourseByID(pkg.CourseID)
		if err != nil {
			return nil, 0, err
		}

		packageType, err := u.packageTypeRepo.GetPackageTypeByID(pkg.PackageTypeID)
		if err != nil {
			return nil, 0, err
		}

		packageTypeData := make(map[string]interface{})
		packageTypeData["id"] = packageType.ID
		packageTypeData["title"] = packageType.Title
		eachPackage.PackageType = packageTypeData

		courseData := make(map[string]interface{})
		courseData["id"] = course.ID
		courseData["courseID"] = course.CourseID
		eachPackage.Course = courseData

		if pkg.TestSeriesID != 0 {
			testSeries, err := u.testSeriesRepo.GetTestSeriesByID(pkg.TestSeriesID)
			if err != nil {
				return nil, 0, err
			}

			testSeriesData := make(map[string]interface{})
			testSeriesData["id"] = testSeries.ID
			testSeriesData["title"] = testSeries.Title
			eachPackage.TestSeries = testSeriesData
		}

		if pkg.LiveGroupID != 0 {
			liveGroup, err := u.liveGroupRepo.GetLiveGroupByID(pkg.LiveGroupID)
			if err != nil {
				return nil, 0, err
			}

			liveGroupData := make(map[string]interface{})
			liveGroupData["id"] = liveGroup.ID
			liveGroupData["title"] = liveGroup.Title
			eachPackage.LiveGroup = liveGroupData
		}

		if pkg.LiveID != 0 {
			live, err := u.liveRepo.GetLiveByID(pkg.LiveID)
			if err != nil {
				return nil, 0, err
			}

			liveData := make(map[string]interface{})
			liveData["id"] = live.ID
			liveData["topic"] = live.Topic
			eachPackage.LiveGroup = liveData
		}

		allPackages = append(allPackages, eachPackage)
	}

	return allPackages, int(totalPage), nil
}
