package gorm

import (
	"avyaas/internal/domain/models"
	"sync"
)

func (repo *Repository) AssignContentsToRelation(relationID, subjectID uint, contentIDs []uint) error {

	var unitChapterContent models.UnitChapterContent

	transaction := repo.db.Begin()

	err := repo.db.Model(&models.UnitChapterContent{}).Where("id  = ?", relationID).First(&unitChapterContent).Error

	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for idx, contentID := range contentIDs {

		wg.Add(1)

		go func(contentID uint) {

			defer wg.Done()

			if unitChapterContent.ContentID == 0 && idx == 0 {
				err := transaction.Model(&models.UnitChapterContent{}).Where("id = ?", relationID).Updates(&models.UnitChapterContent{ContentID: contentID}).Error

				if err != nil {
					transaction.Rollback()
					return
				}

			} else {

				var tempUnitChapterContent models.UnitChapterContent

				err := transaction.Where("content_id  = ? and unit_id  = ? and chapter_id  = ?", contentID, unitChapterContent.UnitID, unitChapterContent.ChapterID).
					FirstOrCreate(&models.UnitChapterContent{
						ContentID: contentID,
						UnitID:    unitChapterContent.UnitID,
						ChapterID: unitChapterContent.ChapterID,
					}).Scan(&tempUnitChapterContent).Error

				if err != nil {
					transaction.Rollback()
					return
				}

				err = transaction.Where("subject_id  = ? and unit_chapter_content_id = ?", subjectID, tempUnitChapterContent.ID).
					FirstOrCreate(&models.SubjectUnitChapterContent{
						SubjectID:            subjectID,
						UnitChapterContentID: tempUnitChapterContent.ID,
					}).Error

				if err != nil {
					transaction.Rollback()
					return
				}
			}

		}(contentID)

	}

	wg.Wait()

	transaction.Commit()

	return nil
}
