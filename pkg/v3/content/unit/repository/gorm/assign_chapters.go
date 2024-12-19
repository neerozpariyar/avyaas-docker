package gorm

import (
	"avyaas/internal/domain/models"
)

func (repo *Repository) AssignChaptersToUnit(subjectID, relationID uint, chapterIDs []uint) error {

	var unitChapterContent models.UnitChapterContent

	transaction := repo.db.Begin()

	if err := repo.db.Model(&models.UnitChapterContent{}).Where("id  = ?", relationID).First(&unitChapterContent).Error; err != nil {
		return err
	}

	for idx, chapterID := range chapterIDs {

		if unitChapterContent.ChapterID == 0 && idx == 0 {

			err := transaction.Model(&models.UnitChapterContent{}).Where("id  = ?", relationID).Update("chapter_id", chapterID).Error

			if err != nil {
				transaction.Rollback()
				return err
			}

		} else {

			var tempSubjectRelation models.UnitChapterContent

			err := transaction.Model(&models.UnitChapterContent{}).
				Where("unit_id  = ? and chapter_id  = ?", unitChapterContent.UnitID, chapterID).
				FirstOrCreate(&models.UnitChapterContent{
					UnitID:    unitChapterContent.UnitID,
					ChapterID: chapterID,
				}).Scan(&tempSubjectRelation).Error

			if err != nil {
				transaction.Rollback()
				return err
			}

			err = transaction.Model(&models.SubjectUnitChapterContent{}).Where("subject_id  = ? and unit_chapter_content_id  = ?", subjectID, tempSubjectRelation.ID).FirstOrCreate(&models.SubjectUnitChapterContent{
				SubjectID:            subjectID,
				UnitChapterContentID: tempSubjectRelation.ID,
			}).Error

			if err != nil {
				transaction.Rollback()
				return err
			}

		}
	}

	transaction.Commit()

	return nil
}
