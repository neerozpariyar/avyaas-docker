package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"
	"time"
)

func (repo *Repository) UpdateTestSeries(data presenter.TestSeriesCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	testSeries, err := repo.GetTestSeriesByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Parse and set the string type start data to *time.Time if provided
	var st time.Time
	if data.StartDate != "" {
		startDate := data.StartDate
		if st, err = time.Parse(time.RFC3339, startDate); err != nil {
			errMap["startDate"] = fmt.Errorf("error parsing invalid UTC time").Error()
			return errMap
		}
	}

	updatedTestSeries := &models.TestSeries{
		Title:       data.Title,
		Description: data.Description,
		NoOfTests:   data.NoOfTests,
		CourseID:    data.CourseID,
		StartDate:   &st,
	}

	transaction := repo.db.Begin()

	// delete the package associated with test series
	if *testSeries.IsPackage && !data.IsPackage {
		var tsPackage *models.Package
		err = repo.db.Where("test_series_id = ? AND test_id = ? AND live_group_id = ? AND live_id = ?", testSeries.ID, 0, 0, 0).First(&tsPackage).Delete(&models.Package{}).Error

		// err = transaction.Unscoped().Where("test_series_id = ? AND package_type_id = ?", data.ID, ).Delete(&models.Package{}).Error
		if err != nil {
			errMap["error"] = err.Error()
			transaction.Rollback()
			return errMap
		}

		err = transaction.Debug().Omit("created_at").Save(&models.TestSeries{
			Timestamp: models.Timestamp{
				ID: data.ID,
			},
			Title:       data.Title,
			Description: data.Description,
			NoOfTests:   data.NoOfTests,
			CourseID:    data.CourseID,
			StartDate:   &st,
			IsPackage:   &data.IsPackage,
			Price:       0,
			Period:      0,
		}).Error
		if err != nil {
			errMap["error"] = err.Error()
			transaction.Rollback()
			return errMap
		}
	}

	// create a new package associated with test series
	if !*testSeries.IsPackage && data.IsPackage {
		err := transaction.Create(&models.Package{
			Title:         data.Title,
			PackageTypeID: data.PackageTypeID,
			CourseID:      data.CourseID,
			TestSeriesID:  data.ID,
			Price:         data.Price,
			Period:        data.Period,
		}).Error
		if err != nil {
			errMap["error"] = err.Error()
			transaction.Rollback()
			return errMap
		}

		updatedTestSeries.IsPackage = &data.IsPackage
		updatedTestSeries.Price = data.Price
		updatedTestSeries.Period = data.Period
	}

	err = transaction.Where("id = ?", data.ID).Updates(&updatedTestSeries).Error
	if err != nil {
		errMap["error"] = err.Error()
		transaction.Rollback()
		return errMap
	}

	transaction.Commit()
	return nil
}
