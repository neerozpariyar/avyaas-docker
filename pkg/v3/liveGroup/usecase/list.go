package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (u *usecase) ListLiveGroup(request presenter.ListLiveGroupRequest) ([]presenter.LiveGroupListResponse, int, error) {
	liveGroups, totalPage, err := u.repo.ListLiveGroup(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allLiveGroups []presenter.LiveGroupListResponse
	for _, liveGroup := range liveGroups {
		singleLiveGroup := presenter.LiveGroupListResponse{
			ID:          liveGroup.ID,
			Title:       liveGroup.Title,
			Description: liveGroup.Description,
			// IsPremium:   liveGroup.IsPremium,
			// Amount:      liveGroup.Amount,
			// ParticipantLimit: liveGroup.ParticipantLimit,
		}

		course, err := u.courseRepo.GetCourseByID(liveGroup.CourseID)
		if err != nil {
			return nil, 0, err
		}

		courseData := make(map[string]interface{})
		courseData["id"] = course.ID
		courseData["courseID"] = course.CourseID
		singleLiveGroup.Course = courseData

		allLiveGroups = append(allLiveGroups, singleLiveGroup)
	}

	return allLiveGroups, int(totalPage), nil
}
