package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type BlogUsecase interface {
	CreateBlog(data presenter.BlogCreateUpdatePresenter) (*models.Blog, error)
	ListBlog(res presenter.BlogListReq) ([]presenter.BlogListRes, int, error)
	UpdateBlog(requestBody presenter.BlogCreateUpdatePresenter) (*models.Blog, map[string]string)
	DeleteBlog(id uint) (*models.Blog, error)
	DetailsOfBlog(id uint) (*presenter.BlogDetailsPresenter, error)
	BlogLikeUnlike(uID, bID uint) (*models.Blog, error)
	CreateBlogComment(data models.BlogComment) (*models.Blog, error)
	ListComments(res presenter.BlogCommentListReq) ([]presenter.BlogCommentListRes, int, error)
	UpdateBlogComment(requestBody models.BlogComment) (*models.BlogComment, map[string]string)
	DeleteBlogComment(id uint) (*models.BlogComment, error)
}

type BlogRepository interface {
	CreateBlog(data presenter.BlogCreateUpdatePresenter) (*models.Blog, error)
	GetBlogByID(id uint) (*models.Blog, error)
	ListBlog(requestBody presenter.BlogListReq) ([]models.Blog, float64, error)
	UpdateBlog(requestBody presenter.BlogCreateUpdatePresenter) error
	DeleteBlog(id uint) (*models.Blog, error)
	DetailsOfBlog(id uint) (*models.Blog, error)
	BlogLikeUnlike(uID, bID uint) (*models.Blog, error)
	CreateBlogComment(data models.BlogComment) (*models.Blog, error)
	GetCommentByID(id uint) (*models.BlogComment, error)
	ListComments(requestBody presenter.BlogCommentListReq) ([]models.BlogComment, float64, error)
	UpdateBlogComment(requestBody models.BlogComment) error
	DeleteBlogComment(id uint) (*models.BlogComment, error)
	GetCommentByBlogID(id uint) (*models.BlogComment, error)
}
