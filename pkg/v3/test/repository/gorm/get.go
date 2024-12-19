package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/*
GetTestTypeByID is a repository method responsible for retrieving a test type from the database
based on its unique identifier (ID).

Parameters:
  - id: A uint representing the unique identifier (ID) of the test type to be retrieved.

Returns:
  - testType: A models.TestType instance representing the retrieved test type.
  - error:    An error, if any, encountered during the database retrieval operation.
*/
func (repo *Repository) GetTestTypeByID(id uint) (models.TestType, error) {
	var testType models.TestType

	// Retrieve the test type from the database based on given id
	err := repo.db.Where("id = ?", id).First(&testType).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.TestType{}, fmt.Errorf("test type with id: '%d' not found", id)
		}

		return models.TestType{}, err
	}

	return testType, nil
}

/*
GetTestTypeByName is a repository method responsible for retrieving a test type from the
database based on its title.

Parameters:
  - title: A string representing the title of the test type to be retrieved.

Returns:
  - testType: A models.TestType instance representing the retrieved test type.
  - error:    An error, if any, encountered during the database retrieval operation.
*/
func (repo *Repository) GetTestTypeByName(title string) (models.TestType, error) {
	var testType models.TestType

	// Retrieve the TestType from the database based on given title
	err := repo.db.Where("title = ?", title).First(&testType).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.TestType{}, err
		}

		return models.TestType{}, err
	}

	return testType, nil
}

/*
GetTestByID is a repository method responsible for retrieving a test from the database based on its
unique identifier (ID).

Parameters:
  - id: A uint representing the unique identifier (ID) of the test to be retrieved.

Returns:
  - test: A models.Test instance representing the retrieved test.
  - error:    An error, if any, encountered during the database retrieval operation.
*/
func (repo *Repository) GetTestByID(id uint) (models.Test, error) {
	var test models.Test

	// Retrieve the test from the database based on given id
	err := repo.db.Where("id = ?", id).Preload("QuestionSets").First(&test).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Test{}, fmt.Errorf("test with id: '%d' not found", id)
		}

		return models.Test{}, err
	}

	return test, nil
}

func (repo *Repository) GetStudentTest(userID, testID uint) (*models.StudentTest, error) {
	var test *models.StudentTest

	err := repo.db.Where("user_id = ? AND test_id = ?", userID, testID).First(&test).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return test, nil
}

func (repo *Repository) GetStudentTestResult(userID, testID uint) (*models.TestResult, error) {
	var result *models.TestResult

	err := repo.db.Where("user_id = ? AND test_id = ?", userID, testID).First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("association for student id:'%d' with test id: '%d' not found", userID, testID)
		}

		return nil, err
	}

	return result, nil
}

func (repo *Repository) GetTestQuestionSet(testID, questionSetID uint) (*models.Test, error) {
	var test *models.Test

	err := repo.db.Where("test_id = ? AND question_set_id = ?", testID, questionSetID).First(&test).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("association for test id:'%d' with question_set id: '%d' not found", testID, questionSetID)
		}

		return nil, err
	}

	return test, nil
}

func (repo *Repository) GetTestsByTestSeriesID(testSeriesID uint) ([]models.Test, error) {
	var tests []models.Test

	err := repo.db.Where("test_series_id = ?", testSeriesID).Find(&tests).Error

	if err != nil {
		return nil, err
	}

	return tests, nil
}

func (repo *Repository) GetTestAttempt(testID, userID uint) (uint, error) {
	var attempt uint

	err := repo.db.Select("attempt").Where("test_id = ? and user_id=?", testID, userID).Find(&attempt).Error

	if err != nil {
		return 0, err
	}

	return attempt, nil
}
