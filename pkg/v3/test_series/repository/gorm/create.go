package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"
	"time"
)

func (repo *Repository) CreateTestSeries(data presenter.TestSeriesCreateUpdateRequest) map[string]string {
	errMap := make(map[string]string)
	var testSeries *models.TestSeries
	var err error

	// Parse and set the string type start data to *time.Time if provided
	var st time.Time
	if data.StartDate != "" {
		startDate := data.StartDate
		if st, err = time.Parse(time.RFC3339, startDate); err != nil {
			errMap["startDate"] = fmt.Errorf("error parsing invalid UTC time").Error()
			return errMap
		}
	}

	// Create a new TestSeries without the ID field from the request
	testSeries = &models.TestSeries{
		Title:       data.Title,
		Description: data.Description,
		NoOfTests:   data.NoOfTests,
		CourseID:    data.CourseID,
		StartDate:   &st,
	}

	if data.IsPackage {
		testSeries.IsPackage = &data.IsPackage
		testSeries.Price = data.Price
		testSeries.Period = data.Period
	}

	transaction := repo.db.Begin()

	err = transaction.Create(&testSeries).Error
	if err != nil {
		errMap["error"] = err.Error()
		transaction.Rollback()
		return errMap
	}

	if data.IsPackage {
		err := transaction.Create(&models.Package{
			Title:         data.Title,
			PackageTypeID: data.PackageTypeID,
			CourseID:      data.CourseID,
			TestSeriesID:  testSeries.ID,
			Price:         data.Price,
			Period:        data.Period,
		}).Error
		if err != nil {
			transaction.Rollback()
			errMap["error"] = err.Error()
			return errMap
		}
	}

	transaction.Commit()
	return errMap
}
