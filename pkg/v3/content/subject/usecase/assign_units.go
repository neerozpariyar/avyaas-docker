package usecase

import "fmt"

func (uCase *usecase) AssignUnitsToSubject(subjectID uint, unitIds []uint) map[string]string {
	errMap := make(map[string]string)

	if _, err := uCase.repo.GetSubjectByID(subjectID); err != nil {

		errMap["Subject"] = fmt.Sprintf("Subject  %d does not Exist", subjectID)

		return errMap

	}

	for _, unitId := range unitIds {

		if _, err := uCase.unitRepo.GetUnitByID(unitId); err != nil {

			errMap["unit"] = fmt.Sprintf("Unit  %d does not Exist", unitId)

			return errMap
		}
	}

	err := uCase.repo.AssignUnitsToSubject([]uint{subjectID}, unitIds)

	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return nil
}
