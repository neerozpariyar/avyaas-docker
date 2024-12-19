package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"math"
)

/*
ListQuestion retrieves a paginated list of question from the database.

Parameters:
  - request: A ListQuestionRequest presenter struct that contains page, courseID, subjectID and
    questionSetID.

Returns:
  - questions: A slice of models.Question representing the retrieved questions.
  - totalPage: A floating-point number representing the total number of pages available.
  - err: An error indicating the success or failure of the operation.
*/
func (repo *Repository) ListTypeQuestion(request presenter.ListQuestionRequest) ([]models.TypeQuestion, float64, error) {
	var err error
	var questions []models.TypeQuestion
	baseQuery := repo.db.Debug().Model(&models.TypeQuestion{})

	if request.Search != "" {
		baseQuery = baseQuery.Where("title LIKE ?", "%"+request.Search+"%")
	}

	if request.SubjectID != 0 {
		baseQuery = baseQuery.Where("subject_id = ?", request.SubjectID)
	}

	if request.QuestionSetID != 0 {
		baseQuery = baseQuery.Joins("JOIN question_set_questions ON question_set_questions.type_question_id = type_questions.id").
			Where("question_set_questions.question_set_id = ? AND type_questions.case_question_id IS NULL", request.QuestionSetID)
	}

	// Calculate the total number of pages based on the given page size
	totalRecords := int64(0)
	baseQuery.Count(&totalRecords)
	totalPage := math.Ceil(float64(totalRecords) / float64(request.PageSize))

	if totalPage < float64(request.Page) {
		request.Page = int(totalPage)
	}

	// Apply pagination
	baseQuery = baseQuery.Offset((request.Page - 1) * request.PageSize).Limit(request.PageSize)

	// Execute the query
	err = baseQuery.Find(&questions).Error
	if err != nil {
		return nil, 0, err
	}

	return questions, totalPage, nil
}
