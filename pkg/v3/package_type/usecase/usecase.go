package usecase

import "avyaas/internal/domain/interfaces"

type usecase struct {
	repo        interfaces.PackageTypeRepository
	serviceRepo interfaces.ServiceRepository
}

func New(repo interfaces.PackageTypeRepository, serviceRepo interfaces.ServiceRepository) *usecase {
	return &usecase{
		repo:        repo,
		serviceRepo: serviceRepo,
	}
}
