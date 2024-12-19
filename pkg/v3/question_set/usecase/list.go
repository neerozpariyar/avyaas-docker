package usecase

import (
	"avyaas/internal/domain/presenter"
)

/*
ListQuestionSet retrieves a paginated list of question sets from the repository.

Parameters:
  - page: An integer representing the page number for pagination.
  - courseID: A uint representing the courseID to filter the question sets.

Returns:
  - questionSets: A slice of models.QuestionSet struct representing the retrieved question sets.
  - totalPage: An integer representing the total number of pages available.
  - error: An error indicating the success or failure of the operation.
*/
func (u *usecase) ListQuestionSet(page int, courseID uint, pageSize int) ([]presenter.QuestionSetDetailsPresenter, int, error) {
	var response []presenter.QuestionSetDetailsPresenter

	// Delegate the retrieval of question sets
	questionSets, totalPage, err := u.repo.ListQuestionSet(page, courseID, pageSize)
	if err != nil {
		return nil, int(totalPage), err
	}

	for _, questionSet := range questionSets {
		course, err := u.courseRepo.GetCourseByID(questionSet.CourseID)
		if err != nil {
			return response, 0, err
		}

		courseData := make(map[string]interface{})
		courseData["id"] = course.ID
		courseData["courseID"] = course.CourseID

		questionSetPresenter := presenter.QuestionSetDetailsPresenter{
			ID:             questionSet.ID,
			Title:          questionSet.Title,
			Description:    questionSet.Description,
			TotalQuestions: questionSet.TotalQuestions,
			Marks:          questionSet.Marks,
			Course:         courseData,
			File:           questionSet.File,
		}

		response = append(response, questionSetPresenter)
	}

	return response, int(totalPage), nil
}
