package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (u *usecase) ListCourse(request presenter.CourseListRequest) ([]presenter.CourseResponse, int, error) {
	cm, totalPage, err := u.repo.ListCourse(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var courses []presenter.CourseResponse

	for i := range cm {

		var courseGroupsForCourse []presenter.CourseGroupForCourse

		courseGroups, err := u.repo.GetCourseGroupByCourseID(cm[i].ID)
		if err != nil {
			return nil, int(totalPage), err
		}

		for _, courseGroup := range courseGroups {
			singleCourseGroupData := presenter.CourseGroupForCourse{
				ID:            courseGroup.ID,
				Title:         courseGroup.Title,
				CourseGroupId: courseGroup.GroupID,
			}

			courseGroupsForCourse = append(courseGroupsForCourse, singleCourseGroupData)

		}
		course := presenter.CourseResponse{
			ID:           cm[i].ID,
			CourseID:     cm[i].CourseID,
			Title:        cm[i].Title,
			Description:  cm[i].Description,
			Available:    cm[i].Available,
			CourseGroups: courseGroupsForCourse,
		}

		if cm[i].Thumbnail != "" {
			course.Thumbnail = utils.GetFileURL(cm[i].Thumbnail)
		}

		courses = append(courses, course)
	}

	return courses, int(totalPage), nil
}
