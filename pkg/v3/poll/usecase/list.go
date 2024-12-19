package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (u *usecase) ListPoll(request presenter.PollListRequest) ([]presenter.Poll, int, error) {
	pm, totalPage, err := u.repo.ListPoll(request)

	if err != nil {
		return nil, int(totalPage), err
	}

	var polls []presenter.Poll
	for i := range pm {

		var allCourses []presenter.CourseDataForPoll

		user, err := u.accountRepo.GetUserByID(request.UserID)
		if err != nil {
			return nil, int(totalPage), err
		}

		subject, err := u.subjectRepo.GetSubjectByID(pm[i].SubjectID)
		if err != nil {
			return nil, int(totalPage), err
		}

		courses, err := u.subjectRepo.GetCoursesBySubjectId(subject.ID)
		if err != nil {
			return nil, int(totalPage), err
		}

		votedOption, err := u.repo.GetVotedOptionByUserID(request.UserID, pm[i].ID)
		if err != nil {
			return nil, int(totalPage), err
		}

		userData := make(map[string]interface{})
		userData["id"] = pm[i].UserID
		userData["name"] = user.FirstName + " " + user.LastName

		for _, course := range courses {
			singleCourseData := presenter.CourseDataForPoll{
				ID:       course.ID,
				Title:    course.Title,
				CourseID: course.CourseID,
			}

			allCourses = append(allCourses, singleCourseData)
		}

		subjectData := make(map[string]interface{})
		subjectData["id"] = pm[i].SubjectID
		subjectData["title"] = subject.Title

		optionData := make([]map[string]interface{}, len(pm[i].Options))

		for j := range pm[i].Options {
			voteCount, err := u.repo.GetVoteCountForOption(pm[i].ID, pm[i].Options[j].Option)
			if err != nil {
				return nil, int(totalPage), err
			}
			optionData[j] = make(map[string]interface{})
			optionData[j]["id"] = pm[i].Options[j].ID
			optionData[j]["title"] = pm[i].Options[j].Option
			optionData[j]["voteCount"] = voteCount
		}

		polls = append(polls, presenter.Poll{
			ID:        pm[i].ID,
			CreatedAt: pm[i].CreatedAt.UTC().Format("2006-01-02T15:04:05Z"),
			Courses:   allCourses,
			Subject:   subjectData,
			// TotalVotes: voteCount,
			VotedOption: votedOption,
			Options:     optionData,
			Question:    pm[i].Question,
			CreatedBy:   userData,
		})
	}
	return polls, int(totalPage), nil
}
