package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) GetTestLeaderboard(testID uint) (*presenter.LeaderboardResponse, error) {
	results, err := uCase.repo.GetTestLeaderboard(testID)
	if err != nil {
		return nil, err
	}

	test, err := uCase.repo.GetTestByID(testID)
	if err != nil {
		return nil, err
	}

	testData := make(map[string]interface{})
	testData["id"] = test.ID
	testData["title"] = test.Title

	course, err := uCase.courseRepo.GetCourseByID(test.CourseID)
	if err != nil {
		return nil, err
	}

	courseData := make(map[string]interface{})
	courseData["id"] = course.ID
	courseData["courseID"] = course.CourseID

	leaderboard := &presenter.LeaderboardResponse{
		Test:   testData,
		Course: courseData,
	}

	var userResponses []presenter.LeaderboardUserResponse
	for _, result := range results {
		user, err := uCase.accountRepo.GetUserByID(result.UserID)
		if err != nil {
			return nil, err
		}

		student, err := uCase.accountRepo.GetStudentByID(result.UserID)
		if err != nil {
			return nil, err
		}

		userResponse := presenter.LeaderboardUserResponse{
			ID:                 result.UserID,
			Name:               user.FirstName + " " + user.MiddleName + " " + user.LastName,
			Email:              user.Email,
			Phone:              user.Phone,
			Score:              result.Score,
			Percentage:         result.Percentage,
			TotalCorrect:       result.TotalCorrect,
			TotalWrong:         result.TotalWrong,
			RegistrationNumber: student.RegistrationNumber,
			TotalAttempted:     result.TotalAttempted,
			TotalUnattempted:   result.TotalUnattempted,
		}

		userResponses = append(userResponses, userResponse)
	}

	leaderboard.Users = userResponses

	return leaderboard, nil
}
