package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
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
func (repo *Repository) ListQuestion(request presenter.ListQuestionRequest) ([]models.Question, float64, error) {
	var err error
	var questions []models.Question
	// query := repo.db.Model(&models.Question{}).Scopes(utils.Paginate(request.Page))
	baseQuery := repo.db.Debug().Model(&models.Question{})

	if request.Search != "" {
		conditionData := make(map[string]interface{})
		conditionData["title"] = "%" + request.Search + "%"

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Question{}, conditionData, false, []string{"like"}, repo.db, request.PageSize)

		if totalPage < float64(request.Page) {
			request.Page = int(totalPage)
		}

		err := baseQuery.Where("title like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&questions).Error

		return questions, totalPage, err
	}

	// Retrieve the questions based on the courseID provided
	// if request.CourseID != 0 {
	// 	// Initialize an empty map to store condition data
	// 	conditionData := make(map[string]interface{})
	// 	conditionData["course_id"] = request.CourseID

	// 	var subjectIDs []string
	// 	err = baseQuery.Select("id").Where("course_id = ?", request.CourseID).Scan(&subjectIDs).Error
	// 	if err != nil {
	// 		return nil, 0, fmt.Errorf("failed to retrieve subjects for courseID: '%d'", request.CourseID)
	// 	}

	// 	// Calculate the total number of pages based on the configured page size for given filter condition
	// 	totalPage := utils.GetTotalPageByConditionModel(models.Test{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

	// 	// query = query.Where("subject_id IN ?", subjectIDs)
	// 	err = baseQuery.Where("subject_id IN ?", subjectIDs).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&questions).Error
	// 	if err != nil {
	// 		return nil, 0, fmt.Errorf("failed to retrieve questions for subjectID: '%d'", request.SubjectID)
	// 	}

	// 	return questions, totalPage, err
	// }

	// Retrieve the questions based on the subjectID provided
	if request.SubjectID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["subject_id"] = request.SubjectID
		if request.Search != "" {
			conditionData := make(map[string]interface{})
			conditionData["title"] = "%" + request.Search + "%"

			// Calculate the total number of pages based on the configured page size for given filter condition
			totalPage := utils.GetTotalPageByConditionModel(models.Question{}, conditionData, false, []string{"like", "like"}, repo.db, request.PageSize)

			if totalPage < float64(request.Page) {
				request.Page = int(totalPage)
			}

			err := baseQuery.Where("title like ?", "%"+request.Search+"%").Scopes(utils.Paginate(request.Page, request.PageSize)).Order("id").Find(&questions).Error

			return questions, totalPage, err
		}
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Question{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		// Fetch a paginated list of questions from the database
		err := baseQuery.Where("subject_id = ?", request.SubjectID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&questions).Error

		return questions, totalPage, err
	}

	// Retrieve the questions based on the questionSetID provided
	if request.QuestionSetID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["question_set_id"] = request.QuestionSetID
		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.Question{}, conditionData, true, []string{"="}, repo.db, request.PageSize)

		// query = query.Where("question_set_id = ?", request.QuestionSetID)
		// Fetch a paginated list of questions from the database
		err := baseQuery.Where("question_set_id = ?", request.QuestionSetID).Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&questions).Error

		return questions, totalPage, err
	}

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.Question{}, repo.db, request.PageSize)

	// Fetch a paginated list of tests from the database
	// err = repo.db.Debug().Model(&models.Question{}).Scopes(utils.Paginate(request.Page)).Find(&questions).Order("id").Error
	err = baseQuery.Scopes(utils.Paginate(request.Page, request.PageSize)).Find(&questions).Error //Where("subject_id = ?", request.SubjectID).
	// err = query.Find(&questions).Order("position").Error
	return questions, totalPage, err
}
