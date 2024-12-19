package usecase

import (
	"avyaas/internal/domain/models"
)

func (uCase *usecase) AssignContentsToChapter(data models.ChapterContent) error {

	// Check if the content already belongs to any chapter with the same cID
	_, err := uCase.repo.GetCourseIDByContentID(data.ContentID)
	if err != nil {
		return err
	}

	// Check if data.ChapterID belongs to cID
	// chapter, err := uCase.chapterRepo.GetChapterByID(data.ChapterID)
	// if err != nil {
	// 	return err
	// }

	// if chapter.Unit.Subject.CourseID == cID {
	// 	return fmt.Errorf("content is already assigned to another chapter in the same course")
	// }

	if err := uCase.repo.AssignContentsToChapter(data); err != nil {
		return err
	}

	return nil
}
