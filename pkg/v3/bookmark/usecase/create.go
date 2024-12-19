package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) CreateBookmark(data presenter.BookmarkCreateUpdateRequest) map[string]string {
	var err error
	errMap := make(map[string]string)
	if data.ContentID != 0 {
		if _, err := uCase.contentRepo.GetContentByID(data.ContentID); err != nil {
			errMap["content_id"] = err.Error()
			return errMap
		}

		if _, err := uCase.contentRepo.CheckStudentContent(data.UserID, data.ContentID); err != nil {
			errMap["student_content_id"] = err.Error()
			return errMap
		}
	}
	if data.QuestionID != 0 {
		if _, err := uCase.questionRepo.GetQuestionByID(data.QuestionID); err != nil {
			errMap["question_id"] = err.Error()
			return errMap
		}

	}
	if err = uCase.repo.CreateBookmark(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
