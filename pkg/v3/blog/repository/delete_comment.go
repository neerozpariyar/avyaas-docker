package repository

import "avyaas/internal/domain/models"

func (repo *Repository) DeleteBlogComment(id uint) (*models.BlogComment, error) {
	var err error

	_, err = repo.GetCommentByID(id)

	if err != nil {
		return nil, err
	}

	err = repo.db.Debug().Where("id = ?", id).Delete(&models.BlogComment{}).Error

	if err != nil {

		return nil, err
	}

	return nil, err
}
