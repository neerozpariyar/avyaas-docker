package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"
)

func (uCase *usecase) CreateReply(data presenter.ReplyCreateUpdateRequest) map[string]string {
	var err error

	errMap := make(map[string]string)

	if _, err := uCase.courseRepo.GetCourseByID(data.CourseID); err != nil {
		errMap["courseID"] = err.Error()
		return errMap
	}
	var discussion models.Discussion
	if discussion, err = uCase.discussionRepo.GetDiscussionByID(data.DiscussionID); err != nil {
		errMap["discussionID"] = err.Error()
		return errMap
	}
	if discussion.CourseID != data.CourseID {
		errMap["discussionID"] = fmt.Sprintf("Discussion ID %d does not belong to Course ID %d", data.DiscussionID, data.CourseID)
		return errMap
	}
	if err = uCase.repo.CreateReply(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap

}
