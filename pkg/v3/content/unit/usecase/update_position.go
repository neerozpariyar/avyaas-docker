package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) UpdateUnitPosition(data presenter.UpdateUnitPositionRequest) map[string]string {
	return uCase.repo.UpdateUnitPosition(data)
}
