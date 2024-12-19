package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (u *usecase) ListUnit(page int, subjectID uint, search string, pageSize int) ([]presenter.Unit, int, error) {
	um, totalPage, err := u.repo.ListUnit(page, subjectID, search, pageSize)
	if err != nil {
		return nil, int(totalPage), err
	}

	var units []presenter.Unit

	for i := range um {
		unit := presenter.Unit{
			ID:          um[i].ID,
			Title:       um[i].Title,
			Description: um[i].Description,
		}

		// subject, err := u.subjectRepo.GetSubjectByID(um[i].SubjectID)
		// if err != nil {
		// 	return nil, int(totalPage), err
		// }

		// subjectData := map[string]interface{}{
		// 	"id":        subject.ID,
		// 	"title":     subject.Title,
		// 	"subjectID": subject.SubjectID,
		// }

		// unit.Subject = subjectData

		// course, err := u.courseRepo.GetCourseByID(subject.CourseID)
		// if err != nil {
		// 	return nil, 0, err
		// }

		// courseData := make(map[string]interface{})
		// courseData["id"] = course.ID
		// courseData["courseID"] = course.CourseID
		// unit.Course = courseData

		if um[i].Thumbnail != "" {
			unit.Thumbnail = utils.GetFileURL(um[i].Thumbnail)
		}

		units = append(units, unit)
	}

	return units, int(totalPage), nil
}
