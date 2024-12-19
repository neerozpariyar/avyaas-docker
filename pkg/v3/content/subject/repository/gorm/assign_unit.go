package gorm

import "avyaas/internal/domain/models"

func (repo *Repository) AssignUnitsToSubject(subjectIds, unitIds []uint) error {

	var totalAssignedSubjects int64

	for _, subjectID := range subjectIds {
		for _, unitID := range unitIds {
			var unitChapterContent models.UnitChapterContent

			err := repo.db.Model(&models.SubjectUnitChapterContent{}).Where("subject_id =?", subjectID).Count(&totalAssignedSubjects).Error

			if err != nil {
				return err
			}

			err = repo.db.Create(&models.UnitChapterContent{
				UnitID: unitID,
			}).Scan(&unitChapterContent).Error

			if err != nil {
				return err
			}

			err = repo.db.Exec("INSERT INTO subject_unit_chapter_contents (subject_id, unit_chapter_content_id, position) VALUES(?,?,?)", subjectID, unitChapterContent.ID, totalAssignedSubjects+1).Error

			if err != nil {
				return err
			}
		}

	}
	return nil
}
