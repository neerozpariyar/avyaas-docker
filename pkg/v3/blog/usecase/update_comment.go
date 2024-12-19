package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) UpdateBlogComment(requestBody models.BlogComment) (*models.BlogComment, map[string]string) {
	var err error

	errMap := make(map[string]string)

	comment, err := uCase.repo.GetCommentByID(requestBody.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	err = uCase.repo.UpdateBlogComment(requestBody)
	if err != nil {
		errMap["update"] = err.Error()
		return nil, errMap
	}
	return comment, errMap
}
