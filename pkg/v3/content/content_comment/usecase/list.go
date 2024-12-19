package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (u *usecase) ListComment(request presenter.CommentListRequest) ([]presenter.CommentListResponse, int, error) {
	cm, totalPage, err := u.repo.ListComment(request)
	if err != nil {
		return nil, int(totalPage), err
	}
	// Update the "updated_at" value for each comment in the list
	var comments []presenter.CommentListResponse

	for i := range cm {
		comment := presenter.CommentListResponse{
			UpdatedAt: cm[i].UpdatedAt,
			Comment:   cm[i].Comment,
			ContentID: cm[i].ContentID,
			ID:        cm[i].ID,
		}

		user, err := u.accountRepo.GetUserByID(cm[i].UserID)
		if err != nil {
			return comments, 0, err
		}

		userData := make(map[string]interface{})
		userData["id"] = user.ID
		userData["name"] = user.FirstName + " " + user.LastName

		comment.CreatedBy = userData
		comments = append(comments, comment)

	}
	return comments, int(totalPage), nil
}
