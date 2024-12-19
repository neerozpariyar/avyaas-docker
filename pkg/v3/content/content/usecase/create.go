package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"strings"
)

func (uCase *usecase) CreateContent(data presenter.ContentCreateUpdateRequest) map[string]string {
	var err error

	errMap := make(map[string]string)

	// if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
	// 	errMap["courseID"] = err.Error()
	// 	return errMap
	// }

	// if _, err := uCase.chapterRepo.GetChapterByID(data.ChapterID); err != nil {
	// 	errMap["chapterID"] = err.Error()
	// 	return errMap
	// }

	contentTypes := []string{"VIDEO", "PDF"}
	if !utils.Contains(contentTypes, strings.ToUpper(string(data.ContentType))) {
		errMap["contentType"] = "invalid content type"
		return errMap
	}

	if err = uCase.repo.CreateContent(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
