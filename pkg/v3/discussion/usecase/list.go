package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (u *usecase) ListDiscussion(request presenter.DiscussionListRequest) ([]presenter.Discussion, int, error) {
	dm, totalPage, err := u.repo.ListDiscussion(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var discussions []presenter.Discussion
	for i := range dm {
		discussion := presenter.Discussion{
			ID:         dm[i].ID,
			Title:      dm[i].Title,
			Query:      dm[i].Query,
			ReplyCount: dm[i].ReplyCount,
			VoteCount:  dm[i].VoteCount,
			Views:      dm[i].Views,
		}

		subject, err := u.subjectRepo.GetSubjectByID(dm[i].SubjectID)
		if err != nil {
			return nil, int(totalPage), err
		}

		subjectData := make(map[string]interface{})
		subjectData["id"] = subject.ID
		subjectData["title"] = subject.Title

		discussion.Subject = subjectData

		course, err := u.courseRepo.GetCourseByID(dm[i].CourseID)
		if err != nil {
			return nil, int(totalPage), err
		}

		courseData := make(map[string]interface{})
		courseData["id"] = dm[i].CourseID
		courseData["courseID"] = course.CourseID

		discussion.Course = courseData

		user, err := u.accountRepo.GetUserByID(request.UserID)
		if err != nil {
			return discussions, 0, err
		}

		userData := make(map[string]interface{})
		userData["id"] = user.ID
		userData["name"] = user.FirstName + " " + user.LastName
		discussion.CreatedBy = userData

		if user.RoleID == 4 { //append the hasLiked value only if the user is a student
			hasLiked, err := u.repo.GetHasLikedValue(dm[i].ID, request.UserID)
			if err != nil {
				return nil, 0, err
			}
			discussion.HasLiked = hasLiked
		}

		discussions = append(discussions, discussion)
	}

	return discussions, int(totalPage), nil
}
