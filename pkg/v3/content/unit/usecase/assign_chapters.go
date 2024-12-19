package usecase

import (
	"fmt"
)

func (uCase *usecase) AssignChaptersToUnit(subjectID uint, relationID uint, chapterIDs []uint) map[string]string {
	errMap := make(map[string]string)

	if _, err := uCase.subjectRepo.GetSubjectByID(subjectID); err != nil {

		errMap["Unit"] = fmt.Sprintf("Subject  %d does not Exist", subjectID)

		return errMap

	}

	for _, chapterID := range chapterIDs {

		if _, err := uCase.chapterRepo.GetChapterByID(chapterID); err != nil {

			errMap["chapter"] = fmt.Sprintf("chapter  %d does not Exist", chapterID)

			return errMap
		}
	}

	err := uCase.repo.AssignChaptersToUnit(subjectID, relationID, chapterIDs)

	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return nil
}
