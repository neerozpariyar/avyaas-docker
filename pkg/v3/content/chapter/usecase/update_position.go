package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) UpdateChapterPosition(data presenter.UpdateChapterPositionRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Delegate the update of content
	if errMap = uCase.repo.UpdateChapterPosition(data); len(errMap) != 0 {
		errMap["error"] = err.Error()
	}

	return errMap
}
