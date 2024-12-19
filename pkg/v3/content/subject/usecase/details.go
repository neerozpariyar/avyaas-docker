package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) GetSubjectDetails(userID, subjectID uint) (presenter.SubjectDetailResponse, error) {
	_, err := uCase.repo.GetSubjectByID(subjectID)
	if err != nil {
		return presenter.SubjectDetailResponse{}, err
	}

	courses, err := uCase.repo.GetCoursesBySubjectId(subjectID)

	if err != nil {
		return presenter.SubjectDetailResponse{}, err
	}

	user, err := uCase.accountRepo.GetUserByID(userID)
	if err != nil {
		return presenter.SubjectDetailResponse{}, err
	}

	if user.RoleID != 1 {
		for _, course := range courses {
			if _, err := uCase.courseRepo.CheckStudentCourse(userID, course.ID); err != nil {
				return presenter.SubjectDetailResponse{}, err
			}

		}
	}

	subjectDetail, err := uCase.repo.GetSubjectDetails(subjectID, userID)
	if err != nil {
		return presenter.SubjectDetailResponse{}, err
	}

	// subjectDetailsResponse := &presenter.SubjectDetailsPresenter{
	// 	ID:          subject.ID,
	// 	SubjectID:   subject.SubjectID,
	// 	Title:       subject.Title,
	// 	Description: subject.Description,
	// 	Thumbnail:   subject.Thumbnail,
	// }

	// for _, unit := range subjectDetails.Units {
	// 	unitPresenter := presenter.UnitDetailsPresenter{
	// 		ID:          unit.ID,
	// 		Title:       unit.Title,
	// 		Description: unit.Description,
	// 		Thumbnail:   unit.Thumbnail,
	// 	}

	// 	// for _, chapter := range unit.Chapters {
	// 	// 	chapterPresenter := presenter.ChapterDetailsPresenter{
	// 	// 		ID:    chapter.ID,
	// 	// 		Title: chapter.Title,
	// 	// 	}

	// 	// 	for _, content := range chapter.Contents {
	// 	// 		contentPresenter := presenter.ContentDetailsPresenter{
	// 	// 			ID:          content.ID,
	// 	// 			Title:       content.Title,
	// 	// 			IsPremium:   content.IsPremium,
	// 	// 			ContentType: content.ContentType,
	// 	// 			Length:      content.Length,
	// 	// 		}
	// 	// 		if user.RoleID == 4 {
	// 	// 			studentContent, err := uCase.contentRepo.CheckStudentContent(user.ID, content.ID)
	// 	// 			if err != nil {
	// 	// 				return nil, err
	// 	// 			}

	// 	// 			contentPresenter.Paid = studentContent.Paid
	// 	// 		}
	// 	// 		sContent, err := uCase.contentRepo.GetContentProgressByContentID(content.ID, user.ID)
	// 	// 		if err != nil {
	// 	// 			if !errors.Is(err, gorm.ErrRecordNotFound) {

	// 	// 				return nil, err
	// 	// 			}
	// 	// 		}
	// 	// 		contentPresenter.Progress = sContent.Progress
	// 	// 		contentPresenter.HasCompleted = *sContent.HasCompleted
	// 	// 		chapterPresenter.Content = append(chapterPresenter.Content, contentPresenter)
	// 	// 	}

	// 	// 	unitPresenter.Chapter = append(unitPresenter.Chapter, chapterPresenter)
	// 	// }

	// 	subjectDetailsResponse.Unit = append(subjectDetailsResponse.Unit, unitPresenter)
	// }

	return subjectDetail, nil
}
