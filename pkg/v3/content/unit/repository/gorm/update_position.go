package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"
)

func (repo *Repository) UpdateUnitPosition(data presenter.UpdateUnitPositionRequest) map[string]string {
	errMap := make(map[string]string)
	transaction := repo.db.Begin()
	fmt.Printf("data.UnitIDs: %v\n", data.UnitIDs)

	for idx, unitID := range data.UnitIDs {
		_, err := repo.GetUnitByID(unitID)
		if err != nil {
			errMap["unitID"] = err.Error()
		}

		err = transaction.Model(&models.Unit{}).Where("id = ?", unitID).Update("position", idx+1).Error
		if err != nil {
			errMap["unitID"] = err.Error()
		}
	}

	if len(errMap) != 0 {
		transaction.Rollback()
		return errMap
	}

	transaction.Commit()
	return errMap
}
