package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

/*
ListTeacher is a usecase function responsible for retrieving a paginated list of teachers.

Parameters:
  - uCase: A pointer to the use case struct, representing the business logic for user-related
    operations. It is used to access the repository for retrieving the list of teachers.
  - page: An integer representing the page number of the paginated result set to be retrieved.

Returns:
  - []models.User: An array of User models representing the paginated list of teachers.
  - int: The total number of pages in the paginated result set.
  - error: An error indicating any issues encountered during the retrieval of the teacher list.
    A nil error signifies a successful retrieval.
*/
func (uCase *usecase) ListTeacher(request *presenter.TeacherListRequest) ([]presenter.TeacherListResponse, int, error) {
	// Delegate the retrieval of teachers
	teachers, totalPage, err := uCase.repo.ListTeacher(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allTeachers []presenter.TeacherListResponse

	for i := range teachers {
		data, err := uCase.repo.GetTeacherByID(teachers[i].ID)
		if err != nil {
			return allTeachers, int(totalPage), err
		}

		singleTeacher := presenter.TeacherListResponse{
			ID:            teachers[i].ID,
			FirstName:     teachers[i].FirstName,
			MiddleName:    teachers[i].MiddleName,
			LastName:      teachers[i].LastName,
			Email:         teachers[i].Email,
			Phone:         teachers[i].Phone,
			Gender:        teachers[i].Gender,
			ReferralCode:  data.ReferralCode,
			ReferralCount: uint(data.ReferralCount),
		}

		// course, err := uCase.courseRepo.GetCourseByID(teacher.CourseID)
		// if err != nil {
		// 	return nil, 0, err
		// }

		// courseData := make(map[string]interface{})
		// courseData["id"] = course.ID
		// courseData["courseID"] = course.CourseID
		// singleTeacher.Course = courseData

		// subject, err := uCase.subjectRepo.GetSubjectByID(teacher.SubjectID)
		// if err != nil {
		// 	return nil, 0, err
		// }

		// subjectData := make(map[string]interface{})
		// subjectData["id"] = subject.ID
		// subjectData["title"] = subject.Title
		// singleTeacher.Subject = subjectData

		if teachers[i].Image != "" {
			singleTeacher.Image = utils.GetFileURL(teachers[i].Image)
		}

		allTeachers = append(allTeachers, singleTeacher)
	}

	return allTeachers, int(totalPage), nil
}
