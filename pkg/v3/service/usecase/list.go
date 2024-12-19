package usecase

import (
	"avyaas/internal/domain/models"
)

func (uCase *usecase) ListService(page int, search string, pageSize int) ([]models.Service, int, error) {
	services, totalPage, err := uCase.repo.ListService(page, search, pageSize)
	if err != nil {
		return nil, int(totalPage), err
	}

	return services, int(totalPage), nil
}
