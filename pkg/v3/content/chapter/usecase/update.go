package usecase

import (
	"avyaas/internal/domain/models"
	"fmt"
)

func (uCase *usecase) UpdateChapter(chapter models.Chapter) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing chapter  with the provided chapter 's ID
	chap, err := uCase.repo.GetChapterByID(chapter.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	chapByID, err := uCase.repo.GetChapterByID(chapter.ID)
	if err == nil {
		// Check if the chapterID is the same as of the requested chapter
		if chap.ID != chapByID.ID {
			errMap["chapterID"] = fmt.Errorf("chapter  with  id: '%v' already exists", chapByID.ID).Error()
			return errMap
		}
	}

	// Delegate the update of chapter
	if err = uCase.repo.UpdateChapter(chapter); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
