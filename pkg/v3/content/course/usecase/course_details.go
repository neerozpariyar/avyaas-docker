package usecase

import "avyaas/internal/domain/presenter"

func (uCase *usecase) GetCourseDetails(userID, courseID uint) (*presenter.CourseDetailsPresenter, error) {
	if _, err := uCase.repo.GetCourseByID(courseID); err != nil {
		return nil, err
	}

	if _, err := uCase.repo.CheckStudentCourse(userID, courseID); err != nil {
		return nil, err
	}

	courseDetails, err := uCase.repo.GetCourseDetails(courseID)
	if err != nil {
		return nil, err
	}

	user, err := uCase.accountRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	courseDetailResponse := &presenter.CourseDetailsPresenter{}

	for _, subject := range courseDetails.Subjects {
		subjectPresenter := presenter.SubjectDetailsPresenter{
			ID:          subject.ID,
			SubjectID:   subject.SubjectID,
			Title:       subject.Title,
			Description: subject.Description,
			Thumbnail:   subject.Thumbnail,
			// CourseID:    subject.CourseID,
		}
		unitChapterContents, err := uCase.repo.GetSubjectHeirarchy(subject.ID)
		if err != nil {
			return nil, err
		}

		for _, unitChapterContent := range unitChapterContents {

			unitPresenter := presenter.UnitDetailsPresenter{
				ID:          unitChapterContent.Unit.ID,
				Title:       unitChapterContent.Unit.Title,
				Description: unitChapterContent.Unit.Description,
				Thumbnail:   unitChapterContent.Unit.Thumbnail,
				// SubjectID:   unit.SubjectID,
			}

			chapterPresenter := presenter.ChapterDetailsPresenter{
				ID:    unitChapterContent.Chapter.ID,
				Title: unitChapterContent.Chapter.Title,
				// UnitID: chapter.UnitID,
			}

			contentPresenter := presenter.ContentDetailsPresenter{
				ID:          unitChapterContent.Content.ID,
				Title:       unitChapterContent.Content.Title,
				IsPremium:   unitChapterContent.Content.ContentIsPremium,
				ContentType: unitChapterContent.Content.ContentType,
				Length:      unitChapterContent.Content.ContentLength,
			}

			chapterPresenter.Content = append(chapterPresenter.Content, contentPresenter)

			unitPresenter.Chapter = append(unitPresenter.Chapter, chapterPresenter)

			subjectPresenter.Unit = append(subjectPresenter.Unit, unitPresenter)
		}

		courseDetailResponse.Subject = append(courseDetailResponse.Subject, subjectPresenter)
	}

	if user.RoleID == 4 {
		studentCourse, err := uCase.repo.CheckStudentCourse(user.ID, courseID)
		if err != nil {
			return nil, err
		}

		courseDetailResponse.Progress = studentCourse.Progress

		if studentCourse.ExpiryDate != nil {
			courseDetailResponse.ExpiryDate = studentCourse.ExpiryDate.UTC().Format("2006-01-02T15:04:05Z")
		}
	}

	return courseDetailResponse, nil
}
