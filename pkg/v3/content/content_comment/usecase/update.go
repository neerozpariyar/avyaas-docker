package usecase

import "avyaas/internal/domain/models"

func (uCase *usecase) UpdateComment(data models.Comment) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Retrieve the existing comment with the provided comment ID
	existingComment, err := uCase.repo.GetCommentByID(data.ID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	// Check if the user is the owner of the comment
	if existingComment.UserID != data.UserID {
		errMap["error"] = " not allowed to update others comment"
		return errMap
	}

	// Delegate the update of comment
	if err = uCase.repo.UpdateComment(data); err != nil {
		errMap["error"] = err.Error()
		return errMap
	}

	return errMap
}
