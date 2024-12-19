package usecase

import "avyaas/internal/domain/presenter"

func (uCase *usecase) ListComments(res presenter.BlogCommentListReq) ([]presenter.BlogCommentListRes, int, error) {
	blogComments, totalPage, err := uCase.repo.ListComments(res)
	if err != nil {
		return nil, int(totalPage), err
	}
	var allComments []presenter.BlogCommentListRes
	for _, blogComment := range blogComments {
		eachComment := &presenter.BlogCommentListRes{
			Comment: blogComment.Comment,
		}
		if blogComment.UserID != 0 {
			user, err := uCase.accountRepo.GetUserByID(blogComment.UserID)
			if err != nil {
				return nil, 0, err
			}

			userData := make(map[string]interface{})
			userData["id"] = user.ID

			eachComment.User = userData
		}

		if blogComment.BlogID != 0 {
			blog, err := uCase.repo.GetBlogByID(blogComment.BlogID)
			if err != nil {
				return nil, 0, err
			}

			blogData := make(map[string]interface{})
			blogData["id"] = blog.ID

			eachComment.Blog = blogData
		}

		allComments = append(allComments, *eachComment)

	}
	return allComments, int(totalPage), err
}
