package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (u *usecase) ListSubject(page int, courseID uint, search string, pageSize int) ([]presenter.SubjectResponse, int, error) {
	sm, totalPage, err := u.repo.ListSubject(page, courseID, search, pageSize)
	if err != nil {
		return nil, int(totalPage), err
	}

	var subjects []presenter.SubjectResponse

	for i := range sm {

		var allCourses []presenter.CourseForSubject

		courses, err := u.repo.GetCoursesBySubjectId(sm[i].ID)
		if err != nil {
			return nil, int(totalPage), err
		}

		for _, course := range courses {
			singleCourse := presenter.CourseForSubject{
				ID:       course.ID,
				Title:    course.Title,
				CourseID: course.CourseID,
			}

			allCourses = append(allCourses, singleCourse)
		}
		subject := presenter.SubjectResponse{
			ID:          sm[i].ID,
			Courses:     allCourses,
			Title:       sm[i].Title,
			SubjectID:   sm[i].SubjectID,
			Description: sm[i].Description,
		}

		if sm[i].Thumbnail != "" {
			subject.Thumbnail = utils.GetFileURL(sm[i].Thumbnail)
		}

		subjects = append(subjects, subject)
	}

	return nil, int(0), nil
}
