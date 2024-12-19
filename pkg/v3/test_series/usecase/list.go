package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) ListTestSeries(request presenter.ListTestSeriesRequest) ([]presenter.TestSeriesListResponse, int, error) {
	testSeries, totalPage, err := uCase.repo.ListTestSeries(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allTestSeries []presenter.TestSeriesListResponse

	for _, ts := range testSeries {
		eachTS := presenter.TestSeriesListResponse{
			ID:          ts.ID,
			Title:       ts.Title,
			Description: ts.Description,
			NoOfTests:   ts.NoOfTests,
			IsPackage:   *ts.IsPackage,
			Price:       ts.Price,
			Period:      ts.Period,
		}

		course, err := uCase.courseRepo.GetCourseByID(ts.CourseID)
		if err != nil {
			return nil, 0, err
		}

		courseData := make(map[string]interface{})
		courseData["id"] = course.ID
		courseData["courseID"] = course.CourseID
		eachTS.Course = courseData

		if ts.StartDate != nil {
			eachTS.StartDate = ts.StartDate.UTC().Format("2006-01-02T15:04:05Z")
		}

		if *ts.IsPackage {
			pkg, err := uCase.packageRepo.GetPackageByTestSeriesID(ts.ID)
			if err != nil {
				return nil, 0, err
			}

			packageType, err := uCase.packageTypeRepo.GetPackageTypeByID(pkg.PackageTypeID)
			if err != nil {
				return nil, 0, err
			}

			packageTypeData := make(map[string]interface{})
			packageTypeData["id"] = packageType.ID
			packageTypeData["title"] = packageType.Title

			eachTS.PackageType = packageTypeData
		}

		allTestSeries = append(allTestSeries, eachTS)
	}

	return allTestSeries, int(totalPage), nil
}
