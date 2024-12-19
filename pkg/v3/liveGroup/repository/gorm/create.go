package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"

	"fmt"
	"time"
)

func (repo *Repository) CreateLiveGroup(data presenter.LiveGroupCreateUpdatePresenter) map[string]string {
	transaction := repo.db.Begin()
	errMap := make(map[string]string)
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

	liveGroup := &models.LiveGroup{
		Title:       data.Title,
		CourseID:    data.CourseID,
		Description: data.Description,
		StartDate:   &st,
		IsPackage:   &data.IsPackage,
		Price:       data.Price,
		Period:      data.Period,
	}

	err = transaction.Create(&liveGroup).Error
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
			LiveGroupID:   liveGroup.ID,
			Price:         data.Price,
			Period:        data.Period,
		}).Error
		if err != nil {
			errMap["error"] = err.Error()
			transaction.Rollback()
			return errMap
		}
	}

	transaction.Commit()
	return errMap
}
