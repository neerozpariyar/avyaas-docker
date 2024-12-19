package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type TestSeriesUsecase interface {
	CreateTestSeries(data presenter.TestSeriesCreateUpdateRequest) map[string]string
	ListTestSeries(request presenter.ListTestSeriesRequest) ([]presenter.TestSeriesListResponse, int, error)
	UpdateTestSeries(data presenter.TestSeriesCreateUpdateRequest) map[string]string
	DeleteTestSeries(id uint) error
}

type TestSeriesRepository interface {
	GetTestSeriesByID(id uint) (*models.TestSeries, error)
	GetTestSeriesByTitle(title string) (*models.TestSeries, error)
	CreateTestSeries(data presenter.TestSeriesCreateUpdateRequest) map[string]string
	ListTestSeries(request presenter.ListTestSeriesRequest) ([]models.TestSeries, float64, error)
	UpdateTestSeries(data presenter.TestSeriesCreateUpdateRequest) map[string]string
	DeleteTestSeries(id uint) error
}
