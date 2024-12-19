package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (u *usecase) ListEnrolledCourse(userID uint, page int, search string, pageSize int) ([]presenter.CourseResponse, int, error) {
	courses, totalPage, err := u.repo.ListEnrolledCourse(userID, page, search, pageSize)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allCourses []presenter.CourseResponse

	user, err := u.accountRepo.GetUserByID(userID)
	if err != nil {
		return nil, int(totalPage), err
	}

	for i := range courses {

		var courseGroupsForCourse []presenter.CourseGroupForCourse

		courseGroups, err := u.repo.GetCourseGroupByCourseID(courses[i].ID)
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
			ID:           courses[i].ID,
			CourseID:     courses[i].CourseID,
			Title:        courses[i].Title,
			Description:  courses[i].Description,
			Available:    courses[i].Available,
			CourseGroups: courseGroupsForCourse,
		}

		if courses[i].Thumbnail != "" {
			course.Thumbnail = utils.GetFileURL(courses[i].Thumbnail)
		}

		if user.RoleID == 4 {
			studentCourse, err := u.repo.CheckStudentCourse(user.ID, courses[i].ID)
			if err != nil {
				return nil, int(totalPage), err
			}
			course.Progress = studentCourse.Progress

			if studentCourse.ExpiryDate != nil {
				course.ExpiryDate = studentCourse.ExpiryDate.UTC().Format("2006-01-02T15:04:05Z")
			}
		}

		allCourses = append(allCourses, course)
	}

	return allCourses, int(totalPage), nil
}
