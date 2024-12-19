package usecase

import "fmt"

func (uCase *usecase) AssignContentsToRelation(relationID uint, subjectId uint, contentIDs []uint) map[string]string {
	errMap := make(map[string]string)

	if _, err := uCase.subjectRepo.GetSubjectByID(subjectId); err != nil {

		errMap["subject"] = fmt.Sprintf("Subject  %d does not Exist", subjectId)

		return errMap

	}

	for _, contentID := range contentIDs {

		if _, err := uCase.repo.GetContentByID(contentID); err != nil {

			errMap["content"] = fmt.Sprintf("Content  %d does not Exist", contentID)

			return errMap
		}
	}

	err := uCase.repo.AssignContentsToRelation(relationID, subjectId, contentIDs)

	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return nil
}
