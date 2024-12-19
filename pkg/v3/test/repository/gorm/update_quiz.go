package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

/*
UpdateTest is a repository function that updates the provided test in the database.

Parameters:
  - test: A pointer to the models.Test structure containing information about the test to be updated.

Returns:
  - error: An error, if any, encountered during the database operation. Returns nil on success.
*/
func (repo *Repository) UpdateTest(data presenter.CreateUpdateTestRequest) map[string]string {
	var err error
	var test *models.Test
	errMap := make(map[string]string)
	transaction := repo.db.Begin()

	// Parse and set the string type start time to *time.Time if provided
	var st time.Time
	if data.StartTime != "" {
		startTime := data.StartTime
		if st, err = time.Parse(time.RFC3339, startTime); err != nil {
			errMap["startTime"] = "error parsing invalid UTC time"
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

	// Convert the data to JSON format
	bData, err := json.Marshal(data)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Unmarshal the JSON data into the models.Test structure
	if err = json.Unmarshal(bData, &test); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	test.StartTime = &st
	test.EndTime = &et

	test.ExtraTime = data.ExtraTime

	if data.IsPublic {
		test.Status = "Scheduled"
	} else {
		test.Status = "Inactive"
	}

	err = transaction.Where("id = ?", data.ID).Omit("created_at").Save(&test).Error
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	if data.QuestionSetID != 0 {
		var testQS *models.TestQuestionSet
		err = repo.db.Where("test_id = ?", data.ID).First(&testQS).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = repo.CreateTestQuestionSet(test.ID, data.QuestionSetID, transaction)
				if err != nil {
					transaction.Rollback()
					errMap["error"] = err.Error()
					return errMap
				}
			} else {
				errMap["error"] = err.Error()
				return errMap
			}
		} else {
			err = repo.UpdateTestQuestionSet(data.ID, data.QuestionSetID, transaction)
			if err != nil {
				transaction.Rollback()
				errMap["error"] = err.Error()
				return errMap
			}
		}
	}

	if data.QuestionSetID == 0 {
		err = repo.DeleteTestQuestionSetByTestID(data.ID, transaction)
		if err != nil {
			transaction.Rollback()
			errMap["error"] = err.Error()
			return errMap
		}
	}

	transaction.Commit()
	return nil
}

func (repo *Repository) UpdateTestQuestionSet(testID, questionSetID uint, transaction *gorm.DB) error {
	return transaction.Where("test_id = ?", testID).Updates(&models.TestQuestionSet{
		TestID:        testID,
		QuestionSetID: questionSetID,
	}).Error
}
