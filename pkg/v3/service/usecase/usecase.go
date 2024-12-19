package usecase

import (
	"avyaas/internal/domain/interfaces"
)

/*
usecase represents the service usecase, which contains the necessary components for handling service
related business logic. It includes an service repository for data access.
*/
type usecase struct {
	repo interfaces.ServiceRepository
}

/*
New initializes and returns a new instance of the service usecase. It takes a service repository as
parameter. The usecase is responsible for handling business logic related to service.
*/
func New(repo interfaces.ServiceRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
