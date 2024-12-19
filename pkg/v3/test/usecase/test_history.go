package usecase

import "avyaas/internal/domain/presenter"

func (uCase *usecase) GetTestHistory(request presenter.TestHistoryRequest) ([]presenter.TestHistoryResponse, float64, error) {
	var response []presenter.TestHistoryResponse
	results, totalPage, err := uCase.repo.GetTestHistory(request)
	if err != nil {
		return nil, 0, err
	}

	if len(results) == 0 {
		return response, 0, nil
	}

	for idx := range results {
		test, err := uCase.repo.GetTestByID(results[idx].TestID)
		if err != nil {
			return nil, 0, err
		}

		result := &presenter.TestHistoryResponse{
			ID:               test.ID,
			Title:            test.Title,
			TotalQuestions:   test.QuestionSets[0].TotalQuestions,
			Marks:            test.QuestionSets[0].Marks,
			Score:            int(results[idx].Score),
			TotalAttempted:   results[idx].TotalAttempted,
			TotalUnattempted: results[idx].TotalUnattempted,
			TotalCorrect:     results[idx].TotalCorrect,
			TotalWrong:       results[idx].TotalWrong,
		}

		course, err := uCase.courseRepo.GetCourseByID(test.CourseID)
		if err != nil {
			return nil, 0, err
		}

		result.Course = course.Title

		response = append(response, *result)
	}

	return response, totalPage, nil
}
