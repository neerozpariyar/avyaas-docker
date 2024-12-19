package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (u *usecase) ListBookmark(data presenter.BookmarkListRequest) ([]presenter.BookmarkListResponse, int, error) {
	bm, totalPage, err := u.repo.ListBookmark(data)
	if err != nil {
		return nil, int(totalPage), err
	}

	var bookmarks []presenter.BookmarkListResponse

	for i := range bm {
		// Fill the title field from respective content or question
		title, err := u.repo.GetBookmarkTitleByID(bm[i].ID)
		if err != nil {
			return nil, int(totalPage), err
		}
		bookmarks = append(bookmarks, presenter.BookmarkListResponse{
			ID:           bm[i].ID,
			Title:        title,
			ContentID:    bm[i].ContentID,
			QuestionID:   bm[i].QuestionID,
			UserID:       bm[i].UserID,
			BookmarkType: bm[i].BookmarkType,
		})
	}

	return bookmarks, int(totalPage), nil
}
