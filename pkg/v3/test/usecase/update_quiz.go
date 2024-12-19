package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"
)

/*
UpdateTest is a usecase method responsible for updaing a test instance.

Parameters:
  - test: A test.CreateUpdateTestRequest presenter struct representing the test to be updated.

Returns:
  - errMap: A map[string]string containing error messages, if any, encountered during the process.
*/
func (uCase *usecase) UpdateTest(data presenter.CreateUpdateTestRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Check if the test with the given ID exists
	test, err := uCase.repo.GetTestByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	if _, err := uCase.repo.GetTestTypeByID(uint(data.TestTypeID)); err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

	if data.TestSeriesID != 0 && data.TestSeriesID != test.TestSeriesID {
		var testSeries *models.TestSeries
		testSeries, err := uCase.testSeriesRepo.GetTestSeriesByID(uint(data.TestTypeID))
		if err != nil {
			errMap["courseID"] = err.Error()
			return errMap
		}

		seriesTests, err := uCase.repo.GetTestsByTestSeriesID(testSeries.ID)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}

		if len(seriesTests) == testSeries.NoOfTests {
			errMap["error"] = fmt.Errorf("maximum number of test already assigned for test series: %s", testSeries.Title).Error()
			return errMap
		}
	}

	if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}

	if data.QuestionSetID != 0 {
		if _, err := uCase.questionSetRepo.GetQuestionSetByID(data.QuestionSetID); err != nil {
			errMap["questionSetID"] = err.Error()
			return errMap
		}
	}

	// Call the repository to update the test
	if errMap = uCase.repo.UpdateTest(data); errMap != nil {
		return errMap
	}

	return errMap
}
