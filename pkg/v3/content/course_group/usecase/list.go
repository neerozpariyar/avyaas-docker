package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"encoding/json"
)

/*
ListCourseGroup retrieves a paginated list of course groups from the repository.

Parameters:
  - page: An integer representing the page number for pagination.

Returns:
  - courseGroups: A slice of CourseGroup pointers representing the retrieved course groups.
  - totalPage: An integer representing the total number of pages available.
  - error: An error indicating the success or failure of the operation.
*/
func (u *usecase) ListCourseGroup(page int, search string, pageSize int) ([]presenter.CourseGroupListResponse, int, error) {
	// Delegate the retrieval of course groups
	courseGroups, totalPage, err := u.repo.ListCourseGroup(page, search, pageSize)
	if err != nil {
		return nil, int(totalPage), err
	}

	var response []presenter.CourseGroupListResponse
	for i, courseGroup := range courseGroups {
		var courseGroupResponse presenter.CourseGroupListResponse
		// Convert the test data to JSON format
		bData, err := json.Marshal(courseGroup)
		if err != nil {
			return nil, 0, err
		}

		// Unmarshal the JSON data into the testPresenter.CreateUpdateTestRequest structure
		if err = json.Unmarshal(bData, &courseGroupResponse); err != nil {
			return nil, 0, err
		}

		if courseGroups[i].Thumbnail != "" {
			courseGroupResponse.Thumbnail = utils.GetFileURL(courseGroups[i].Thumbnail)
		}

		courseGroupResponse.NoOfCourses = len(courseGroup.Courses)
		response = append(response, courseGroupResponse)
	}

	return response, int(totalPage), nil
}
