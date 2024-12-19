package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) EvaluateProgress(data presenter.ProgressPresenter) error {
	return uCase.repo.EvaluateProgress(data)
}
