package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) CreateChapter(chapter models.Chapter) map[string]string {
	var err error
	errMap := make(map[string]string)

	if err = uCase.repo.CreateChapter(chapter); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
