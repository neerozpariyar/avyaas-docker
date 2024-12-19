package usecase

import (
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
)

func (uCase *usecase) ListNotice(req presenter.NoticeListReq) ([]presenter.NoticeListPresenter, int, error) {

	notices, totalPage, err := uCase.repo.ListNotice(req)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allNotice []presenter.NoticeListPresenter

	for _, notice := range notices {
		eachNotice := &presenter.NoticeListPresenter{
			ID:          notice.ID,
			Title:       notice.Title,
			Description: notice.Description,
		}
		if notice.File != "" {
			url := utils.GetFileURL(notice.File)
			eachNotice.File = url
		}

		if notice.CourseID != 0 {
			course, err := uCase.courseRepo.GetCourseByID(notice.CourseID)
			if err != nil {
				return nil, 0, err
			}
			courseData := make(map[string]interface{})
			courseData["id"] = course.ID
			courseData["courseID"] = course.CourseID

			eachNotice.Course = courseData
		}

		allNotice = append(allNotice, *eachNotice)

	}
	return allNotice, int(totalPage), err
}
