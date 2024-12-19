package usecase

import (
	"avyaas/internal/domain/presenter"
	"encoding/json"
	"errors"
)

/*
ListTest retrieves a paginated list of tests from the repository.

Parameters:
  - page: An integer representing the page number for pagination.

Returns:
  - allTests: A slice of test.CreateUpdateTestRequest presenter struct representing the retrieved
    tests.
  - totalPage: An integer representing the total number of pages available.
  - error: An error indicating the success or failure of the operation.
*/
// MapTestTypeIDToName maps test type ID to name
// func MapTestTypeIDToName(u *usecase, typeID int) (string, error) {
// 	// Retrieve the test type with the given ID
// 	testType, err := u.repo.GetTestTypeByID(uint(typeID))
// 	if err != nil {
// 		return "", err
// 	}
// 	return testType.Name, nil
// }

func MapTestTypeIDToData(u *usecase, typeID int) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	// Retrieve the test type with the given ID
	testType, err := u.repo.GetTestTypeByID(uint(typeID))
	if err != nil {
		return data, err
	}

	data["id"] = testType.ID
	data["title"] = testType.Title

	return data, nil
}

func (u *usecase) ListTest(request presenter.ListTestRequest) ([]presenter.TestResponse, int, error) {
	// Delegate the retrieval of tests
	tests, totalPage, err := u.repo.ListTest(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	user, err := u.accountRepo.GetUserByID(request.UserID)
	if err != nil {
		return nil, int(totalPage), err
	}

	if user.RoleID == 4 && request.CourseID == 0 {
		return nil, int(totalPage), errors.New("courseID is a required field")
	}

	var allTests []presenter.TestResponse

	for _, eachTest := range tests {
		var test presenter.TestResponse

		// Convert the test data to JSON format
		bData, err := json.Marshal(eachTest)
		if err != nil {
			return nil, 0, err
		}

		// Unmarshal the JSON data into the testPresenter.CreateUpdateTestRequest structure
		if err = json.Unmarshal(bData, &test); err != nil {
			return nil, 0, err
		}

		// Format the start time and end time to UTC string format
		if test.StartTime != "" {
			test.StartTime = eachTest.StartTime.UTC().Format("2006-01-02T15:04:05Z")
		}

		if test.StartTime != "" {
			test.EndTime = eachTest.EndTime.UTC().Format("2006-01-02T15:04:05Z")
		}

		if len(eachTest.QuestionSets) != 0 {
			questionSet, err := u.questionSetRepo.GetQuestionSetByID(eachTest.QuestionSets[0].ID)
			if err != nil {
				return nil, 0, err
			}

			qsData := make(map[string]interface{})
			qsData["id"] = questionSet.ID
			qsData["title"] = questionSet.Title
			test.QuestionSet = qsData

			// Set the total questions and marks of  question set in the test response
			test.TotalQuestions = questionSet.TotalQuestions
			test.Marks = questionSet.Marks
			// test.QuestionSetID = questionSet.ID
		}

		// Set the test type name in the test's type field
		testType, err := MapTestTypeIDToData(u, int(eachTest.TestTypeID))
		if err != nil {
			return nil, 0, err
		}
		// Set the test type title in test's type field
		test.TestType = testType

		course, err := u.courseRepo.GetCourseByID(eachTest.CourseID)
		if err != nil {
			return nil, 0, err
		}

		courseData := make(map[string]interface{})
		courseData["id"] = course.ID
		courseData["courseID"] = course.CourseID
		test.Course = courseData

		if eachTest.TestSeriesID != 0 {
			testSeries, err := u.testSeriesRepo.GetTestSeriesByID(eachTest.TestSeriesID)
			if err != nil {
				return nil, 0, err
			}

			testSeriesData := make(map[string]interface{})
			testSeriesData["id"] = testSeries.ID
			testSeriesData["title"] = testSeries.Title
			test.TestSeries = testSeriesData
		}

		// Add each test data in allTests slice
		allTests = append(allTests, test)
	}

	return allTests, int(totalPage), nil
}
