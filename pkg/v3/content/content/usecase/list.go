package usecase

import (
	"avyaas/internal/domain/presenter"
	"errors"

	"gorm.io/gorm"
)

func (u *usecase) ListContent(request presenter.ContentListRequest) ([]presenter.SingleContentResponse, int, error) {
	contents, totalPage, err := u.repo.ListContent(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	user, err := u.accountRepo.GetUserByID(request.UserID)
	if err != nil {
		return nil, 0, err
	}

	var allContents []presenter.SingleContentResponse
	for _, content := range contents {
		eachContent := presenter.SingleContentResponse{
			ID:          content.ID,
			Title:       content.Title,
			ContentType: content.ContentType,
			IsPremium:   *content.IsPremium,
			Length:      content.Length,
		}
		if user.RoleID == 4 {
			studentContent, err := u.repo.CheckStudentContent(user.ID, content.ID)
			if err != nil {
				return nil, 0, err
			}

			eachContent.Paid = studentContent.Paid
			sContent, err := u.repo.GetContentProgressByContentID(content.ID, user.ID)
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {

					return nil, 0, err
				}
			}
			eachContent.HasCompleted = *sContent.HasCompleted
		}

		allContents = append(allContents, eachContent)
	}

	return allContents, int(totalPage), nil
}
