package usecase

import "avyaas/internal/domain/models"

func (u *usecase) ListFeedback(page int, courseID uint, pageSize int) ([]models.Feedback, int, error) {
	feedbacks, totalPage, err := u.repo.ListFeedback(page, courseID, pageSize)
	if err != nil {
		return nil, int(totalPage), err
	}

	return feedbacks, int(totalPage), nil
}
