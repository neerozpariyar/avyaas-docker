package usecase

import "avyaas/internal/domain/presenter"

func (uCase *usecase) UpdateContentPosition(data presenter.UpdateContentPositionRequest) map[string]string {
	// Delegate the update of content
	return uCase.repo.UpdateContentPosition(data)
}
