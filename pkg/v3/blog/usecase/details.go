package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) DetailsOfBlog(id uint) (*presenter.BlogDetailsPresenter, error) {
	blog, err := uCase.repo.DetailsOfBlog(id)
	if err != nil {
		return nil, err
	}

	comments, err := uCase.repo.GetCommentByBlogID(id)
	if err != nil {
		return nil, err
	}
	detail := &presenter.BlogDetailsPresenter{
		ID:          blog.ID,
		Tags:        blog.Tags,
		Title:       blog.Title,
		Description: blog.Description,
		Views:       blog.Views,
		Cover:       blog.Cover,
		Likes:       blog.Likes,
		Comments:    comments.Comment,
	}

	return detail, nil
}
