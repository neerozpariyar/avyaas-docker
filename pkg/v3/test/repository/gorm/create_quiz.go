package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

/*
CreateTest is a repository function that persists the provided test in the database.

Parameters:
  - test: A pointer to the models.Test structure containing information about the test to be created.

Returns:
  - error: An error, if any, encountered during the database operation. Returns nil on success.
*/
func (repo *Repository) CreateTest(data presenter.CreateUpdateTestRequest) map[string]string {
	transaction := repo.db.Begin()
	var test *models.Test
	var err error
	errMap := make(map[string]string)

	// Parse and set the string type start time to *time.Time if provided
	var st time.Time
	if data.StartTime != "" {
		startTime := data.StartTime
		if st, err = time.Parse(time.RFC3339, startTime); err != nil {
			errMap["startTime"] = "error parsing invald UTC time"
			return errMap
		}
	}

	// Parse and set the string type end time to *time.Time if provided
	var et time.Time
	if data.EndTime != "" {
		endTime := data.EndTime
		if et, err = time.Parse(time.RFC3339, endTime); err != nil {
			errMap["endTime"] = "error parsing invalid UTC time"
			return errMap
		}
	}

	// Marshal the input data to JSON bytes
	bData, err := json.Marshal(data)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Unmarshal JSON bytes to the models.Test structure
	if err = json.Unmarshal(bData, &test); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	test.StartTime = &st
	test.EndTime = &et

	if data.IsPublic {
		test.Status = "Scheduled"
	} else {
		test.Status = "Inactive"
	}

	err = transaction.Create(&test).Error
	if err != nil {
		transaction.Rollback()
		errMap["error"] = err.Error()
		return errMap
	}

	if data.QuestionSetID != 0 {
		err = repo.CreateTestQuestionSet(test.ID, data.QuestionSetID, transaction)
		if err != nil {
			transaction.Rollback()
			errMap["error"] = err.Error()
			return errMap
		}
	}

	// if data.IsPackage{
	if !data.IsPremium && *data.IsPackage {
		err = transaction.Create(&models.Package{
			Title:         data.Title,
			PackageTypeID: 8,
			CourseID:      data.CourseID,
			TestID:        test.ID,
			Price:         data.Price,
		}).Error
		if err != nil {
			transaction.Rollback()
			errMap["error"] = err.Error()
			return errMap
		}
	}

	transaction.Commit()
	return nil
}

func (repo *Repository) CreateTestQuestionSet(testID, questionSetID uint, transaction *gorm.DB) error {
	return transaction.Create(&models.TestQuestionSet{
		TestID:        testID,
		QuestionSetID: questionSetID,
	}).Error
}
