package usecase

import "avyaas/internal/domain/presenter"

func (uCase *usecase) CreateComment(data presenter.CommentCreateUpdateRequest) map[string]string {
	var err error

	errMap := make(map[string]string)

	// if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
	// 	errMap["courseID"] = err.Error()
	// 	return errMap
	// }
	// var content models.Content
	// if content.CourseID != data.CourseID {
	// 	errMap["contentID"] = fmt.Sprintf("Content ID %d does not belong to Course ID %d", data.ContentID, data.CourseID)
	// 	return errMap
	// }
	content, err := uCase.contentRepo.GetContentByID(data.ContentID)
	if err != nil {
		errMap["contentID"] = err.Error()
		return errMap
	}

	if content.ContentType != "VIDEO" {
		errMap["contentID"] = "content ID does not belong to a video; can't comment on it."
		return errMap
	}
	if err = uCase.repo.CreateComment(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
