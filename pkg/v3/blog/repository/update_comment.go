package repository

import "avyaas/internal/domain/models"

func (repo *Repository) UpdateBlogComment(requestBody models.BlogComment) error {
	var err error
	updateComment := &models.BlogComment{
		Comment: requestBody.Comment,
	}

	// if requestBody.BlogID != 0 {
	// 	updateComment.BlogID = requestBody.BlogID
	// }
	// if requestBody.UserID != 0 {
	// 	updateComment.UserID = requestBody.UserID
	// }

	err = repo.db.Where("id = ?", requestBody.ID).Updates(&updateComment).Error
	if err != nil {
		return err
	}

	return err
}
