package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

func (u *usecase) ListLive(request presenter.ListLiveRequest) (interface{}, int, error) {
	user, err := u.accountRepo.GetUserByID(request.UserID)
	if err != nil {
		return "", 0, err
	}

	if user.RoleID == 4 && request.CourseID == 0 {
		return "", 0, errors.New("courseID is a required field")
	}

	lives, totalPage, err := u.repo.ListLive(request)
	if err != nil {
		return "", int(totalPage), err
	}

	var allLives []presenter.LiveListResponse

	for i := range lives {
		var eachLive presenter.LiveListResponse

		bData, err := json.Marshal(lives[i])
		if err != nil {
			return "", 0, err
		}

		err = json.Unmarshal(bData, &eachLive)
		if err != nil {
			return "", 0, err
		}

		if lives[i].LiveGroupID != 0 {
			liveGroup, err := u.liveGroupRepo.GetLiveGroupByID(lives[i].LiveGroupID)
			if err != nil {
				return "", 0, err
			}

			liveGroupData := make(map[string]interface{})
			liveGroupData["id"] = liveGroup.ID
			liveGroupData["title"] = liveGroup.Title
			eachLive.LiveGroup = liveGroupData
		}

		course, err := u.courseRepo.GetCourseByID(lives[i].CourseID)
		if err != nil {
			return "", 0, err
		}

		courseData := make(map[string]interface{})
		courseData["id"] = course.ID
		courseData["courseID"] = course.CourseID
		eachLive.Course = courseData

		subject, err := u.subjectRepo.GetSubjectByID(lives[i].SubjectID)
		if err != nil {
			return "", 0, err
		}

		subjectData := make(map[string]interface{})
		subjectData["id"] = subject.ID
		subjectData["title"] = subject.Title
		eachLive.Subject = subjectData

		if lives[i].StartTime != nil {
			eachLive.StartTime = lives[i].StartTime.UTC().Format("2006-01-02T15:04:05Z")
		}

		if lives[i].EndDateTime != nil {
			eachLive.EndDateTime = lives[i].EndDateTime.UTC().Format("2006-01-02T15:04:05Z")
		}
		allLives = append(allLives, eachLive)
	}

	if user.RoleID == 4 {
		enc, err := utils.Encrypt([]byte(fmt.Sprintf("%v", allLives)))

		if err != nil {
			return "", 0, err
		}

		encryptedLives := base64.URLEncoding.EncodeToString(enc)
		return encryptedLives, int(totalPage), nil
	}

	return allLives, int(totalPage), nil
}
