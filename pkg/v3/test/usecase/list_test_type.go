package usecase

import (
	"avyaas/internal/domain/models"
)

/*
ListTestType retrieves a paginated list of test types from the repository.

Parameters:
  - page: An integer representing the page number for pagination.

Returns:
  - testTypes: A slice of TestType representing the retrieved test types.
  - totalPage: An integer representing the total number of pages available.
  - error: An error indicating the success or failure of the operation.
*/
func (u *usecase) ListTestType(page int, pageSize int) ([]models.TestType, int, error) {
	// Delegate the retrieval of test types
	testTypes, totalPage, err := u.repo.ListTestType(page, pageSize)
	if err != nil {
		return nil, int(totalPage), err
	}

	return testTypes, int(totalPage), nil
}
