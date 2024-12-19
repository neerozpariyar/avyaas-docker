package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"
)

/*
ListQuestionSet retrieves a paginated list of question sets from the database.

Parameters:
  - page: An integer representing the page number for pagination.
  - courseID: A uint representing the courseID to filter the question set.

Returns:
  - data: A slice of models.QuestionSet representing the retrieved question sets.
  - totalPage: A floating-point number representing the total number of pages available.
  - err: An error indicating the success or failure of the operation.
*/
func (repo *Repository) ListQuestionSet(page int, courseID uint, pageSize int) ([]models.QuestionSet, float64, error) {
	var questionSets []models.QuestionSet
	baseQuery := repo.db.Debug().Model(&models.QuestionSet{}).Order("id")

	// Retrieve the question sets based on the courseID provided
	if courseID != 0 {
		// Initialize an empty map to store condition data
		conditionData := make(map[string]interface{})
		conditionData["course_id"] = courseID

		// Calculate the total number of pages based on the configured page size for given filter condition
		totalPage := utils.GetTotalPageByConditionModel(models.QuestionSet{}, conditionData, true, []string{"="}, repo.db, pageSize)

		// Fetch a paginated list of question sets from the database
		err := baseQuery.Where("course_id = ?", courseID).Scopes(utils.Paginate(page, pageSize)).Find(&questionSets).Error

		return questionSets, totalPage, err
	}

	// Calculate the total number of pages based on the configured page size
	totalPage := utils.GetTotalPage(models.QuestionSet{}, repo.db, pageSize)

	// Fetch a paginated list of question sets from the database
	err := baseQuery.Scopes(utils.Paginate(page, pageSize)).Find(&questionSets).Error

	return questionSets, totalPage, err
}
