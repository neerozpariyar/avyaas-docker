package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (u *usecase) ListReply(request presenter.ReplyListRequest) ([]presenter.ReplyListResponse, int, error) {
	rl, totalPage, err := u.repo.ListReply(request)
	if err != nil {
		return nil, int(totalPage), err
	}
	var allReplies []presenter.ReplyListResponse
	// Convert the list of replies to a list of presenter.Reply
	for i := range rl {
		replies := presenter.ReplyListResponse{
			ID:           rl[i].ID,
			DiscussionID: rl[i].DiscussionID,
			CreatedAt:    rl[i].CreatedAt,
			Reply:        rl[i].Reply,
			CourseID:     rl[i].CourseID,
		}
		if rl[i].UserID != 0 {
			user, err := u.accountRepo.GetUserByID(rl[i].UserID)
			if err != nil {
				return nil, int(totalPage), err
			}

			userData := make(map[string]interface{})
			userData["userID"] = user.ID
			userData["username"] = user.Username

			replies.User = userData
		}
		allReplies = append(allReplies, replies)

	}

	return allReplies, int(totalPage), nil
}
