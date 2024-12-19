package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type UnitUsecase interface {
	CreateUnit(data presenter.UnitCreateUpdateRequest) map[string]string
	ListUnit(page int, subjectID uint, search string, pageSize int) ([]presenter.Unit, int, error)
	UpdateUnit(data presenter.UnitCreateUpdateRequest) map[string]string
	DeleteUnit(id uint) error
	AssignChaptersToUnit(uint, uint, []uint) map[string]string
	UpdateUnitPosition(data presenter.UpdateUnitPositionRequest) map[string]string
}

type UnitRepository interface {
	//GetUnitByUnitID(UnitID string) (models.Unit, error)
	GetUnitByID(id uint) (models.Unit, error)

	CreateUnit(data presenter.UnitCreateUpdateRequest) error
	ListUnit(page int, subjectID uint, search string, pageSize int) ([]models.Unit, float64, error)
	UpdateUnit(data presenter.UnitCreateUpdateRequest) error
	DeleteUnit(id uint) error
	AssignChaptersToUnit(uint, uint, []uint) error
	UpdateUnitPosition(data presenter.UpdateUnitPositionRequest) map[string]string
}
